-- +migrate Up
CREATE TRIGGER IF NOT EXISTS update_conversations_after_insert
AFTER INSERT ON messages
BEGIN
    UPDATE conversations SET modified_at = unixepoch() WHERE id = NEW.conversation_id;
    END;