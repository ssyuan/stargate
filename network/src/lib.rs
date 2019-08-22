#![feature(async_await)]

pub mod net;

#[cfg(test)]
mod tests {
    use crate::net::{build_libp2p_service, Message, Service};
    use crypto::ed25519::compat;
    use futures::{future, future::Future, prelude::Async, stream, stream::Stream};
    use libp2p::PeerId;
    use network_libp2p::{build_multiaddr, identity, start_service, NetworkConfiguration};
    use parity_multiaddr::{Multiaddr, Protocol};
    use parking_lot::Mutex;
    use std::time::Duration;
    use std::{sync::Arc, thread};
    use tokio::{runtime::Runtime, sync::mpsc::error::TrySendError};
    use types::account_address::AccountAddress;

    fn build_network_service(num: usize, base_port: u16) -> Vec<Service> {
        let mut result: Vec<Service> = Vec::with_capacity(num);
        let mut first_addr = None::<Multiaddr>;

        for index in 0..num {
            let mut boot_nodes = Vec::new();

            if let Some(first_addr) = first_addr.as_ref() {
                boot_nodes.push(
                    first_addr
                        .clone()
                        .with(Protocol::P2p(
                            result[0].libp2p_service.lock().peer_id().clone().into(),
                        ))
                        .to_string(),
                );
            }

            let config = network_libp2p::NetworkConfiguration {
                listen_addresses: vec![build_multiaddr![
                    Ip4([127, 0, 0, 1]),
                    Tcp(base_port + index as u16)
                ]],
                boot_nodes,
                ..network_libp2p::NetworkConfiguration::default()
            };

            if first_addr.is_none() {
                first_addr = Some(config.listen_addresses.iter().next().unwrap().clone());
            }
            let server = Service::new(config);
            result.push(server);
        }
        result
    }

    #[test]
    fn test_send_receive() {
        let (service1, mut service2) = {
            let mut l = build_network_service(2, 50400).into_iter();
            let a = l.next().unwrap();
            let b = l.next().unwrap();
            (a, b)
        };
        let msg_peer_id = service1.libp2p_service.lock().peer_id().clone();
        let sender_fut = stream::repeat(1)
            .and_then(move |_| {
                match service2.network_sender.try_send(Message {
                    peer_id: msg_peer_id.clone(),
                    msg: vec![1, 2],
                }) {
                    Ok(()) => Ok(Async::Ready(Some(()))),
                    Err(e) => match e.is_full() {
                        true => Ok(Async::NotReady),
                        false => return Err(()),
                    },
                }
            })
            .for_each(|_| Ok(()));

        let rt = Runtime::new().unwrap();
        let executor = rt.executor();
        executor.spawn(sender_fut);
        rt.shutdown_on_idle().wait().unwrap();
    }

    #[test]
    fn test_account_to_peer_id() {
        //let (private_key, public_key) = compat::generate_keypair();
        //let key = private_key.to_bytes();
    }

    #[test]
    fn test_connected_nodes() {
        let (service1, mut service2) = {
            let mut l = build_network_service(2, 50400).into_iter();
            let a = l.next().unwrap();
            let b = l.next().unwrap();
            (a, b)
        };
        thread::sleep(Duration::new(1, 0));
        for (peer_id, peer) in service1.libp2p_service.lock().state().connected_peers {
            println!("id: {:?}, peer: {:?}", peer_id, peer);
        }
    }
}
