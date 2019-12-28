use lazy_static::lazy_static;
use libra_types::account_config::{account_module_name, core_code_address};

use libra_types::{
    identifier::{IdentStr, Identifier},
    language_storage::ModuleId,
    transaction::ScriptAction,
};

lazy_static! {
    static ref CHALLENGE_METHOD_NAME: Identifier = Identifier::new("challenge").unwrap();
    static ref RESOLVE_METHOD_NAME: Identifier = Identifier::new("resolve").unwrap();
    static ref CLOSE_METHOD_NAME: Identifier = Identifier::new("close").unwrap();
}

fn channel_challenge_name() -> &'static IdentStr {
    &*CHALLENGE_METHOD_NAME
}
fn channel_resolve_name() -> &'static IdentStr {
    &*RESOLVE_METHOD_NAME
}
fn channel_close_name() -> &'static IdentStr {
    &*CLOSE_METHOD_NAME
}
pub fn resolve_channel_action() -> ScriptAction {
    ScriptAction::new_call(
        ModuleId::new(core_code_address(), account_module_name().into()),
        channel_resolve_name().into(),
        vec![],
    )
}

pub fn challenge_channel_action() -> ScriptAction {
    ScriptAction::new_call(
        ModuleId::new(core_code_address(), account_module_name().into()),
        channel_challenge_name().into(),
        vec![],
    )
}

pub fn close_channel_action() -> ScriptAction {
    ScriptAction::new_call(
        ModuleId::new(core_code_address(), account_module_name().into()),
        channel_close_name().into(),
        vec![],
    )
}
