-- +migrate Down
-- +migrate StatementBegin
DROP TRIGGER IF EXISTS update_conversations_after_insert;
-- +migrate StatementEnd

-- +migrate Down
-- +migrate StatementBegin
DROP TRIGGER IF EXISTS check_reply_conversation;
-- +migrate StatementEnd