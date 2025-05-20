-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd
CREATE TABLE user_referals ( 
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  referring_user uuid REFERENCES users(id),
  referred_user uuid REFERENCES users(id),
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


CREATE TABLE vibe_point_transactions (
  id uuid PRIMARY KEY DEFAULT gen_random_uuid(),
  awarded_user uuid REFERENCES users(id),
  points_awarded INT NOT NULL,
  reason TEXT NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
  modified_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);


CREATE OR REPLACE FUNCTION award_vibe_points_for_referral()
-- +goose StatementBegin
RETURNS TRIGGER AS 'BEGIN
  INSERT INTO vibe_point_transactions (awarded_user, points_awarded, reason) VALUES (
    NEW.referring_user,
    2,
    ''Referral bonus''
  );

  UPDATE userprofile
  SET vibe_points = vibe_points + 2,
      modified_at = NOW()
  WHERE user_id = NEW.referring_user;
  RETURN NEW;
END
' LANGUAGE plpgsql;
-- +goose StatementEnd


CREATE TRIGGER trigger_award_vibe_points_for_referral
AFTER INSERT ON user_referals
FOR EACH ROW
EXECUTE FUNCTION award_vibe_points_for_referral();


-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';

DROP TRIGGER IF EXISTS trigger_award_vibe_points_for_referral ON user_referals;
DROP FUNCTION IF EXISTS award_vibe_points_for_referral();
DROP TABLE IF EXISTS vibe_point_transactions;
DROP TABLE IF EXISTS user_referals;
-- +goose StatementEnd
