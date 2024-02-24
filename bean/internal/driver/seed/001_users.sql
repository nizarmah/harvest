INSERT INTO users
(id, email)
VALUES
-- statics
('00000000-0000-0000-0000-000000000001', 'user-1'),
('00000000-0000-0000-0000-000000000002', 'user-2'),
-- payment methods related
('00000000-0000-0000-0001-000000000001', 'payment-methods'),
('00000000-0000-0000-0001-000000000002', 'no-payment-methods'),
-- subscriptions related
('00000000-0000-0000-0002-000000000001', 'subscriptions'),
('00000000-0000-0000-0002-000000000002', 'no-subscriptions'),
-- memberships related
('00000000-0000-0000-0003-000000000001', 'active-membership'),
('00000000-0000-0000-0003-000000000002', 'expired-membership'),
('00000000-0000-0000-0003-000000000003', 'no-membership'),

-- filler data
('10000000-0000-0000-0000-000000000001', '277@hey.com')
;

---- create above / drop below ----

DELETE FROM users;
