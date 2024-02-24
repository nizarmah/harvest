INSERT INTO memberships
(user_id, created_at, expires_at)
VALUES
-- test data
-- memberships related
('00000000-0000-0000-0003-000000000001', '2024-01-21 00:00:00', NULL),
('00000000-0000-0000-0003-000000000002', '2024-01-21 00:20:00', '2024-02-21 00:20:00'),

-- filler data
('10000000-0000-0000-0000-000000000001', '2024-01-21 00:00:00', NULL)
;

---- create above / drop below ----

DELETE FROM memberships;
