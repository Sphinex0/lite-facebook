-- +migrate Up
-- +migrate StatementBegin
CREATE TRIGGER IF NOT EXISTS update_conversations_after_insert
AFTER INSERT ON messages
FOR EACH ROW
BEGIN
    -- Update the modified_at timestamp in the conversations table
    UPDATE conversations 
    SET modified_at = unixepoch() 
    WHERE id = NEW.conversation_id;
END;
-- +migrate StatementEnd


-- +migrate Up
-- +migrate StatementBegin
CREATE TRIGGER IF NOT EXISTS check_reply_conversation
BEFORE INSERT ON messages
FOR EACH ROW
WHEN NEW.reply IS NOT NULL
BEGIN
    SELECT RAISE(ABORT, 'Reply must belong to the same conversation')
    WHERE (SELECT conversation_id FROM messages WHERE id = NEW.reply) != NEW.conversation_id;
END;
-- +migrate StatementEnd

