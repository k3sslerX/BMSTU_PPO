INSERT INTO manufacturer (name, country, year_of_foundation, id) VALUES ('Ferrari', 'Italy', 1939, 'f04b6635-1e63-4962-9dfa-78fa3adca11e');
INSERT INTO manufacturer (name, country, year_of_foundation, id) VALUES ('Toyota', 'Japan', 1937, '2d204400-1346-4faa-af6c-5fb1757ef3e9');
INSERT INTO manufacturer (name, country, year_of_foundation, id) VALUES ('Porsche', 'Germany', 1931, '10adf8ee-e8ae-4070-a2e7-1c2b189c10c1');
INSERT INTO manufacturer (name, country, year_of_foundation, id) VALUES ('Cadillac', 'United States', 1902, 'e64fa1c1-b0e7-4a93-94ba-f2a61b762bc3');
INSERT INTO manufacturer (name, country, year_of_foundation, id) VALUES ('BMW', 'Germany', 1916, 'fb6c4e5e-5063-4bae-8a0b-5cc0df8af4ae');
INSERT INTO manufacturer (name, country, year_of_foundation, id) VALUES ('Alpine', 'France', 1955, 'd8bc12e7-3184-406c-9212-1aac2b5b951d');
INSERT INTO manufacturer (name, country, year_of_foundation, id) VALUES ('Peugeot', 'France', 1810, '6f3e08f9-b1f9-46b5-a044-3ba5aff16182');
INSERT INTO manufacturer (name, country, year_of_foundation, id) VALUES ('Aston Martin', 'United Kingdom', 1913, 'a7cd28a4-2d88-45c9-b43f-89f3b034f620');
INSERT INTO manufacturer (name, country, year_of_foundation, id) VALUES ('Lamborghini', 'Italy', 1963, 'b02617d6-54fe-447f-bf74-c6dd466f7634');
INSERT INTO manufacturer (name, country, year_of_foundation, id) VALUES ('Isotta Fraschini', 'Italy', 1900, '71ac33c5-d79e-4c12-af87-d89f89291a13');
INSERT INTO manufacturer (name, country, year_of_foundation, id) VALUES ('Ford', 'United States', 1903, '7aa0897c-0f31-402e-9262-aa71b6b1db39');
INSERT INTO manufacturer (name, country, year_of_foundation, id) VALUES ('Lexus', 'Japan', 1989, 'e3458635-884e-406f-afe5-a51d9f788859');
INSERT INTO manufacturer (name, country, year_of_foundation, id) VALUES ('Corvette', 'United States', 1953, '4320e572-5ac0-4094-b740-d919e4e2a5ef');
INSERT INTO manufacturer (name, country, year_of_foundation, id) VALUES ('Mercedes-AMG', 'Germany', 1967, '012686d8-b319-4a20-ba34-4f6d8628a6fd');
INSERT INTO manufacturer (name, country, year_of_foundation, id) VALUES ('McLaren', 'United Kingdom', 1963, 'a01268a9-358b-4bd1-a3d9-fbebf6b0e4e1');


--
-- Data for Name: raceclass; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO raceclass (name, docs, id) VALUES ('F1', NULL, '5b52b484-ed5a-4805-aeb8-a06f0bcc7dd8');
INSERT INTO raceclass (name, docs, id) VALUES ('GT3', NULL, '5b706798-b242-4d92-ac67-adc241e65d58');
INSERT INTO raceclass (name, docs, id) VALUES ('LMGT3', NULL, 'd264cfb9-f291-4e87-89c0-119120e16d6a');
INSERT INTO raceclass (name, docs, id) VALUES ('HY', NULL, 'e2b39596-4a5b-40d5-800e-ed81d47847e0');


--
-- Data for Name: car; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('296 GT3', 2024, '8be18828-72d8-4bae-b0ed-081a526b8a25', 'f04b6635-1e63-4962-9dfa-78fa3adca11e', '5b706798-b242-4d92-ac67-adc241e65d58');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('296 GT3', 2024, '52c33ff3-3966-4c3f-8b0b-5536d419da30', 'f04b6635-1e63-4962-9dfa-78fa3adca11e', 'd264cfb9-f291-4e87-89c0-119120e16d6a');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('499P', 2022, '8c65d9ba-dc51-42f6-960b-cb02d1f8e630', 'f04b6635-1e63-4962-9dfa-78fa3adca11e', 'e2b39596-4a5b-40d5-800e-ed81d47847e0');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('GR010 Hybrid', 2021, '566afa3e-3087-4d49-863d-832f5f319f27', '2d204400-1346-4faa-af6c-5fb1757ef3e9', 'e2b39596-4a5b-40d5-800e-ed81d47847e0');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('963', 2023, 'd778d138-7529-49cf-b106-74b6d0118b12', '10adf8ee-e8ae-4070-a2e7-1c2b189c10c1', 'e2b39596-4a5b-40d5-800e-ed81d47847e0');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('V-Series.R', 2023, 'aa0d5e92-495f-4eb9-bc16-72a37fed1bb3', 'e64fa1c1-b0e7-4a93-94ba-f2a61b762bc3', 'e2b39596-4a5b-40d5-800e-ed81d47847e0');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('M Hybrid V8', 2023, 'e77e4e08-a380-4514-bc3d-516b204a3379', 'fb6c4e5e-5063-4bae-8a0b-5cc0df8af4ae', 'e2b39596-4a5b-40d5-800e-ed81d47847e0');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('A424', 2024, '8553712f-2a33-44f2-97d1-bef578e0d62d', 'd8bc12e7-3184-406c-9212-1aac2b5b951d', 'e2b39596-4a5b-40d5-800e-ed81d47847e0');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('9X8', 2022, '3245bfa4-ef14-4813-aae5-6f486601d5e7', '6f3e08f9-b1f9-46b5-a044-3ba5aff16182', 'e2b39596-4a5b-40d5-800e-ed81d47847e0');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('Valkyrie AMR-LMH', 2025, '22d8cfcd-ef36-4afb-b35a-f3ac734b0fe6', 'a7cd28a4-2d88-45c9-b43f-89f3b034f620', 'e2b39596-4a5b-40d5-800e-ed81d47847e0');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('SC63', 2024, '23b9bb68-efcd-4146-8eb7-25fe08c6d109', 'b02617d6-54fe-447f-bf74-c6dd466f7634', 'e2b39596-4a5b-40d5-800e-ed81d47847e0');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('Tipo 6 LMH-C', 2024, 'c5f6d54c-b716-4b0e-b8b2-60c4d6fe5485', '71ac33c5-d79e-4c12-af87-d89f89291a13', 'e2b39596-4a5b-40d5-800e-ed81d47847e0');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('911 GT3 R (992)', 2023, '9ea57b4d-ad89-4918-b02c-b11010831d2c', '10adf8ee-e8ae-4070-a2e7-1c2b189c10c1', 'd264cfb9-f291-4e87-89c0-119120e16d6a');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('M4 GT3', 2022, '849cf0fc-b201-46db-a696-ac01a530a956', 'fb6c4e5e-5063-4bae-8a0b-5cc0df8af4ae', 'd264cfb9-f291-4e87-89c0-119120e16d6a');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('720S GT3 Evo', 2023, 'bd17ae64-a499-42f6-9dfa-d20164cb8e13', 'a01268a9-358b-4bd1-a3d9-fbebf6b0e4e1', 'd264cfb9-f291-4e87-89c0-119120e16d6a');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('Vantage AMR GT3 Evo', 2024, '074f072f-5901-40d3-9968-d05d72aa26ab', 'a7cd28a4-2d88-45c9-b43f-89f3b034f620', 'd264cfb9-f291-4e87-89c0-119120e16d6a');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('Huracán GT3 EVO2', 2023, '5c1fafe5-1723-4fcf-8f9e-97b2b35e0752', 'b02617d6-54fe-447f-bf74-c6dd466f7634', 'd264cfb9-f291-4e87-89c0-119120e16d6a');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('Mustang GT3', 2024, '146f126e-9152-4071-8617-28c67bc1beed', '7aa0897c-0f31-402e-9262-aa71b6b1db39', 'd264cfb9-f291-4e87-89c0-119120e16d6a');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('RC F GT3', 2017, '37d5ead9-93e9-4393-b306-d76f1f7a0fe4', 'e3458635-884e-406f-afe5-a51d9f788859', 'd264cfb9-f291-4e87-89c0-119120e16d6a');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('Z06 GT3.R', 2024, '56d8b000-e04b-4331-a4e8-251f95b0ebe1', '4320e572-5ac0-4094-b740-d919e4e2a5ef', 'd264cfb9-f291-4e87-89c0-119120e16d6a');
INSERT INTO car (model, year_of_production, id, manufacturer, raceclass) VALUES ('GT3 Evo', 2020, '7027a0ab-af28-49e2-aba4-f840ec5268d5', '012686d8-b319-4a20-ba34-4f6d8628a6fd', 'd264cfb9-f291-4e87-89c0-119120e16d6a');


--
-- Data for Name: team; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO team (name, logo, country, id) VALUES ('AF Corse', NULL, 'Italy', '49001e74-6d1b-48fb-8ed2-0e7afb9d4153');
INSERT INTO team (name, logo, country, id) VALUES ('Ferrari AF Corse', NULL, 'Italy', 'c88b98d8-148f-401e-99e2-8d5637a5027c');
INSERT INTO team (name, logo, country, id) VALUES ('Toyota Gazoo Racing', NULL, 'Japan', 'd3c9d641-3bb7-4129-9252-1c99e6e3fe68');
INSERT INTO team (name, logo, country, id) VALUES ('Porsche Penske Motorsport', NULL, 'Germany', '50492663-7821-4f8f-b319-f59c19acc9a1');
INSERT INTO team (name, logo, country, id) VALUES ('Proton Competition', NULL, 'Germany', '0177baa4-f617-48d6-8a6e-a5fd43dabeea');
INSERT INTO team (name, logo, country, id) VALUES ('Cadillac Hertz Team JOTA', NULL, 'United Kingdom', '00453069-65a8-4c19-b453-3d7c29b17efe');
INSERT INTO team (name, logo, country, id) VALUES ('BMW M Team WRT', NULL, 'Belgium', 'dd3d5a63-cbe0-406d-a09e-fde2a4f4d5aa');
INSERT INTO team (name, logo, country, id) VALUES ('Alpine Endurance Team', NULL, 'France', '708c8b92-2dfb-413f-8a61-41fa3c209bef');
INSERT INTO team (name, logo, country, id) VALUES ('Peugeot TotalEnergies', NULL, 'France', '8d4aeb49-6e04-4fa7-8e7e-61fb88ae2a32');
INSERT INTO team (name, logo, country, id) VALUES ('Aston Martin THOR Team', NULL, 'United Kingdom', 'b1bc4c43-3b30-4a2d-9243-f5c4581174d5');
INSERT INTO team (name, logo, country, id) VALUES ('Racing Spirit of Léman', NULL, 'France', '49f2ebaa-cc9d-4c19-aaa2-9482a4fcd435');
INSERT INTO team (name, logo, country, id) VALUES ('Vista AF Corse', NULL, 'Italy', 'd9bfa501-e44c-424c-b7c7-5b7c2267129f');
INSERT INTO team (name, logo, country, id) VALUES ('Heart of Racing Team', NULL, 'United Kingdom', '6a8350bb-8059-4a0c-a188-26e476a7a244');
INSERT INTO team (name, logo, country, id) VALUES ('Team WRT', NULL, 'Belgium', '8f1e61a3-66c5-42ac-8a86-528622befc03');
INSERT INTO team (name, logo, country, id) VALUES ('TF Sport', NULL, 'United Kingdom', 'd15fe6dd-9fc2-4e8f-8cd1-0792e9ebc969');
INSERT INTO team (name, logo, country, id) VALUES ('United Autosports', NULL, 'United Kingdom', '9a5abb29-8756-411f-a29b-3aba79b07208');
INSERT INTO team (name, logo, country, id) VALUES ('Iron Lynx', NULL, 'Italy', '507c85bc-77a9-40c3-9c18-773807166289');
INSERT INTO team (name, logo, country, id) VALUES ('Proton Competition', NULL, 'Germany', '6a2524f3-a0be-409c-a491-03c7a4f14ee8');
INSERT INTO team (name, logo, country, id) VALUES ('Akkodis ASP Team', NULL, 'France', 'ee07271f-f0cc-4733-9d46-13f3c8a10aa9');
INSERT INTO team (name, logo, country, id) VALUES ('Manthey 1st Phorm', NULL, 'Germany', '040cc441-c391-4a9c-a62e-4e245319b0c9');
INSERT INTO team (name, logo, country, id) VALUES ('ISOTTA FRASCHINI', NULL, 'Italy', '92ab1623-98f9-418e-a324-d26d0f34aac2');
INSERT INTO team (name, logo, country, id) VALUES ('Manthey EMA (Iron Dames)', NULL, 'Germany', '1cdb506c-a094-4fe6-af46-42fe2840ddbb');


--
-- Data for Name: car_p; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO car_p (number, id, car, team) VALUES ('83', 'b1d55df9-d7b0-4f95-bd08-ba8009ba1b29', '8c65d9ba-dc51-42f6-960b-cb02d1f8e630', '49001e74-6d1b-48fb-8ed2-0e7afb9d4153');
INSERT INTO car_p (number, id, car, team) VALUES ('50', '1e1c1666-982b-46f8-9c4a-4350c0994aac', '8c65d9ba-dc51-42f6-960b-cb02d1f8e630', 'c88b98d8-148f-401e-99e2-8d5637a5027c');
INSERT INTO car_p (number, id, car, team) VALUES ('51', 'f9dc7fe7-bc27-4921-b665-96d1aa72ec87', '8c65d9ba-dc51-42f6-960b-cb02d1f8e630', 'c88b98d8-148f-401e-99e2-8d5637a5027c');
INSERT INTO car_p (number, id, car, team) VALUES ('7', 'e03d5b31-1b44-4375-b09c-a303cbbf5d8f', '566afa3e-3087-4d49-863d-832f5f319f27', 'd3c9d641-3bb7-4129-9252-1c99e6e3fe68');
INSERT INTO car_p (number, id, car, team) VALUES ('8', '00b26904-d63c-45aa-a917-1878c7af767e', '566afa3e-3087-4d49-863d-832f5f319f27', 'd3c9d641-3bb7-4129-9252-1c99e6e3fe68');
INSERT INTO car_p (number, id, car, team) VALUES ('5', '44fff04f-d19a-4aab-8ee4-a7cdcfc38fd8', 'd778d138-7529-49cf-b106-74b6d0118b12', '50492663-7821-4f8f-b319-f59c19acc9a1');
INSERT INTO car_p (number, id, car, team) VALUES ('6', 'c76df0ff-11ca-4170-b88a-1291f8442066', 'd778d138-7529-49cf-b106-74b6d0118b12', '50492663-7821-4f8f-b319-f59c19acc9a1');
INSERT INTO car_p (number, id, car, team) VALUES ('99', '296e1bc6-c1c7-4261-b48a-935986950265', 'd778d138-7529-49cf-b106-74b6d0118b12', '0177baa4-f617-48d6-8a6e-a5fd43dabeea');
INSERT INTO car_p (number, id, car, team) VALUES ('12', 'ddff5320-64f6-460b-a76a-19488ea08e70', 'aa0d5e92-495f-4eb9-bc16-72a37fed1bb3', '00453069-65a8-4c19-b453-3d7c29b17efe');
INSERT INTO car_p (number, id, car, team) VALUES ('38', '147902c8-1f9c-4ed3-a310-91935ee7b41d', 'aa0d5e92-495f-4eb9-bc16-72a37fed1bb3', '00453069-65a8-4c19-b453-3d7c29b17efe');
INSERT INTO car_p (number, id, car, team) VALUES ('15', '3f8291ff-6434-400b-bc28-53a15df6bd18', 'e77e4e08-a380-4514-bc3d-516b204a3379', 'dd3d5a63-cbe0-406d-a09e-fde2a4f4d5aa');
INSERT INTO car_p (number, id, car, team) VALUES ('20', '8793bec9-4be7-459a-b9da-0dd56a223de6', 'e77e4e08-a380-4514-bc3d-516b204a3379', 'dd3d5a63-cbe0-406d-a09e-fde2a4f4d5aa');
INSERT INTO car_p (number, id, car, team) VALUES ('35', '60a41976-bb4d-4ce5-ab9c-da8a09106389', '8553712f-2a33-44f2-97d1-bef578e0d62d', '708c8b92-2dfb-413f-8a61-41fa3c209bef');
INSERT INTO car_p (number, id, car, team) VALUES ('36', '22984b36-d8fd-48f8-8068-006140dd34da', '8553712f-2a33-44f2-97d1-bef578e0d62d', '708c8b92-2dfb-413f-8a61-41fa3c209bef');
INSERT INTO car_p (number, id, car, team) VALUES ('93', 'd8c0d728-08cb-450a-b5ee-3b98027e80d3', '3245bfa4-ef14-4813-aae5-6f486601d5e7', '8d4aeb49-6e04-4fa7-8e7e-61fb88ae2a32');
INSERT INTO car_p (number, id, car, team) VALUES ('94', '011f7d02-b781-443c-bccc-c3d61aab0f24', '3245bfa4-ef14-4813-aae5-6f486601d5e7', '8d4aeb49-6e04-4fa7-8e7e-61fb88ae2a32');
INSERT INTO car_p (number, id, car, team) VALUES ('63', '77ab202c-f44b-4ba6-b947-84276d9c3769', '23b9bb68-efcd-4146-8eb7-25fe08c6d109', '507c85bc-77a9-40c3-9c18-773807166289');
INSERT INTO car_p (number, id, car, team) VALUES ('54', '1655d35e-8a04-47eb-b694-fcc0f8b8c968', '52c33ff3-3966-4c3f-8b0b-5536d419da30', 'd9bfa501-e44c-424c-b7c7-5b7c2267129f');
INSERT INTO car_p (number, id, car, team) VALUES ('21', '624b63a0-41d9-4c3b-af56-d631afebf840', '52c33ff3-3966-4c3f-8b0b-5536d419da30', '49001e74-6d1b-48fb-8ed2-0e7afb9d4153');
INSERT INTO car_p (number, id, car, team) VALUES ('85', 'f334e8a1-28fd-416a-a0ce-dad0423e385b', '9ea57b4d-ad89-4918-b02c-b11010831d2c', '1cdb506c-a094-4fe6-af46-42fe2840ddbb');
INSERT INTO car_p (number, id, car, team) VALUES ('92', '716d317e-bde0-4fa3-9738-ec0ef9c3cce1', '9ea57b4d-ad89-4918-b02c-b11010831d2c', '040cc441-c391-4a9c-a62e-4e245319b0c9');
INSERT INTO car_p (number, id, car, team) VALUES ('77', 'e1450469-8b19-4706-a9cb-5ef41fcf500a', '9ea57b4d-ad89-4918-b02c-b11010831d2c', '0177baa4-f617-48d6-8a6e-a5fd43dabeea');
INSERT INTO car_p (number, id, car, team) VALUES ('88', 'bc43d47e-8913-4931-9042-788772e5a9ba', '9ea57b4d-ad89-4918-b02c-b11010831d2c', '0177baa4-f617-48d6-8a6e-a5fd43dabeea');
INSERT INTO car_p (number, id, car, team) VALUES ('31', 'c65c3f38-77fd-4099-a10d-35d75d6be870', '849cf0fc-b201-46db-a696-ac01a530a956', '8f1e61a3-66c5-42ac-8a86-528622befc03');
INSERT INTO car_p (number, id, car, team) VALUES ('46', '621245bb-31ee-4da9-b224-6e30be870e17', '849cf0fc-b201-46db-a696-ac01a530a956', '8f1e61a3-66c5-42ac-8a86-528622befc03');
INSERT INTO car_p (number, id, car, team) VALUES ('59', '9a46a97f-30ad-43a0-aeda-91a280947848', 'bd17ae64-a499-42f6-9dfa-d20164cb8e13', '9a5abb29-8756-411f-a29b-3aba79b07208');
INSERT INTO car_p (number, id, car, team) VALUES ('95', '1461d107-ea45-4dbe-b456-92cba175a824', 'bd17ae64-a499-42f6-9dfa-d20164cb8e13', '9a5abb29-8756-411f-a29b-3aba79b07208');
INSERT INTO car_p (number, id, car, team) VALUES ('27', '1323c6db-0a94-4b45-bd07-5a7091e777e9', '074f072f-5901-40d3-9968-d05d72aa26ab', '6a8350bb-8059-4a0c-a188-26e476a7a244');
INSERT INTO car_p (number, id, car, team) VALUES ('10', '92116a61-59c8-4b83-ab9e-f4153f3c9337', '074f072f-5901-40d3-9968-d05d72aa26ab', '49f2ebaa-cc9d-4c19-aaa2-9482a4fcd435');
INSERT INTO car_p (number, id, car, team) VALUES ('60', '2ee8a9e4-6224-485d-be98-14b50bd3b608', '5c1fafe5-1723-4fcf-8f9e-97b2b35e0752', '507c85bc-77a9-40c3-9c18-773807166289');
INSERT INTO car_p (number, id, car, team) VALUES ('61', 'dd4f8cdc-eb0a-4e93-94cc-72c958729269', '5c1fafe5-1723-4fcf-8f9e-97b2b35e0752', '507c85bc-77a9-40c3-9c18-773807166289');
INSERT INTO car_p (number, id, car, team) VALUES ('77', '9a97bead-1061-4cc7-85a2-866670c59342', '146f126e-9152-4071-8617-28c67bc1beed', '0177baa4-f617-48d6-8a6e-a5fd43dabeea');
INSERT INTO car_p (number, id, car, team) VALUES ('88', 'd262ce67-5829-4eb6-befa-48d07e08b6c6', '146f126e-9152-4071-8617-28c67bc1beed', '0177baa4-f617-48d6-8a6e-a5fd43dabeea');
INSERT INTO car_p (number, id, car, team) VALUES ('78', '70645b62-8df4-4334-a8d5-79acce411cba', '37d5ead9-93e9-4393-b306-d76f1f7a0fe4', 'ee07271f-f0cc-4733-9d46-13f3c8a10aa9');
INSERT INTO car_p (number, id, car, team) VALUES ('87', 'b5f023e6-d2c9-4e10-841c-536f5eb176eb', '37d5ead9-93e9-4393-b306-d76f1f7a0fe4', 'ee07271f-f0cc-4733-9d46-13f3c8a10aa9');
INSERT INTO car_p (number, id, car, team) VALUES ('33', 'd2335551-2808-47bb-8229-6cd849c3af1a', '56d8b000-e04b-4331-a4e8-251f95b0ebe1', 'd15fe6dd-9fc2-4e8f-8cd1-0792e9ebc969');
INSERT INTO car_p (number, id, car, team) VALUES ('81', 'c2afd8f8-285f-4903-982a-770a2eacfa04', '56d8b000-e04b-4331-a4e8-251f95b0ebe1', 'd15fe6dd-9fc2-4e8f-8cd1-0792e9ebc969');
INSERT INTO car_p (number, id, car, team) VALUES ('60', '28d21e2c-ef12-41e4-ab27-3c17b22c60c1', '7027a0ab-af28-49e2-aba4-f840ec5268d5', '507c85bc-77a9-40c3-9c18-773807166289');
INSERT INTO car_p (number, id, car, team) VALUES ('61', 'f15657ca-919a-420d-8785-f605ce6098ab', '7027a0ab-af28-49e2-aba4-f840ec5268d5', '507c85bc-77a9-40c3-9c18-773807166289');
INSERT INTO car_p (number, id, car, team) VALUES ('11', '628c8f62-fb33-42d8-9337-dfc8c0db0d7c', 'c5f6d54c-b716-4b0e-b8b2-60c4d6fe5485', '92ab1623-98f9-418e-a324-d26d0f34aac2');
INSERT INTO car_p (number, id, car, team) VALUES ('007', 'e9da05f6-4335-493e-8b1a-9d050ce33e59', '22d8cfcd-ef36-4afb-b35a-f3ac734b0fe6', 'b1bc4c43-3b30-4a2d-9243-f5c4581174d5');
INSERT INTO car_p (number, id, car, team) VALUES ('009', '80923a54-81ab-40e7-b9d5-e7da52d83761', '22d8cfcd-ef36-4afb-b35a-f3ac734b0fe6', 'b1bc4c43-3b30-4a2d-9243-f5c4581174d5');


--
-- Data for Name: organizer; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO organizer (name, id) VALUES ('World Endurance Championship', '8fea3615-d820-4fd2-8c5f-bcdc9a61107e');


--
-- Data for Name: championship; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO championship (year, id, organizer) VALUES (2025, '6f79a9bf-5c1a-4e1d-a9ab-304dda6d911f', '8fea3615-d820-4fd2-8c5f-bcdc9a61107e');
INSERT INTO championship (year, id, organizer) VALUES (2024, '097b8e91-3e3c-44a4-adba-8838097e05e5', '8fea3615-d820-4fd2-8c5f-bcdc9a61107e');
INSERT INTO championship (year, id, organizer) VALUES (2023, '3e7ca1d1-4a7d-4a7f-85c0-4f3f70008042', '8fea3615-d820-4fd2-8c5f-bcdc9a61107e');


--
-- Data for Name: driver; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Robert Kubica', NULL, 'Poland', '1984-12-07', '5582f399-3a47-4d7e-b325-59fc7888e364');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Antonio Fuoco', NULL, 'Italy', '1996-05-20', 'd5f10135-5dc3-4b76-9ca8-9cca944e8eb5');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Miguel Molina', NULL, 'Spain', '1989-02-17', 'dd32dd80-8842-450f-9d2e-503e18a2505d');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Nicklas Nielsen', NULL, 'Denmark', '1997-02-06', 'ee95cd16-cd67-40a0-90a0-f975bf91d064');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Yifei Ye', NULL, 'China', '2000-06-16', 'c1f36b6e-47fb-4517-8394-e76021367e50');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Alessandro Pier Guidi', NULL, 'Italy', '1983-12-18', 'a8cf93ca-8be8-4bf7-ab24-e568e44e2bc0');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Antonio Giovinazzi', NULL, 'Italy', '1993-12-14', 'f50c48a9-65e9-4617-bf7a-0e7b0c00b7e3');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Phil Hanson', NULL, 'United Kingdom', '1999-07-05', 'c86a5bb9-2555-4510-8536-5e3cf3fa7cee');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('James Calado', NULL, 'United Kingdom', '1989-06-13', '36d5ae0f-c7ed-4c35-9039-4434e2e5f6ff');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Mike Conway', NULL, 'United Kingdom', '1983-08-19', '59411c83-b6f3-4724-86dc-c54262166a85');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Kamui Kobayashi', NULL, 'Japan', '1986-09-13', '892a168c-9c19-4356-adfe-ec37308c0ae4');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Nyck de Vries', NULL, 'Netherlands', '1995-02-06', '13c8ae17-41d7-4321-898d-e82555325a1d');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Sébastien Buemi', NULL, 'Switzerland', '1988-10-31', '8d3e1f53-0290-414c-b788-fd6d13ce1b8d');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Brendon Hartley', NULL, 'New Zealand', '1989-11-10', '29146774-0472-47a4-8a2e-ead3f93876fc');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Ryo Hirakawa', NULL, 'Japan', '1994-03-07', 'cce9262f-858b-4318-96ea-e9515441d530');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Michael Christensen', NULL, 'Denmark', '1990-08-28', '67484669-106c-4f78-b5f5-6605b464b621');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Julien Andlauer', NULL, 'France', '1999-07-19', '84aa0b68-758d-4a91-8ed8-0a42af389125');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Mathieu Jaminet', NULL, 'France', '1994-10-24', '945da040-e6d2-4544-a36f-a6bf8698b029');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Kévin Estre', NULL, 'France', '1988-10-28', '2a52185e-77b3-40d0-9065-86a959be6552');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Laurens Vanthoor', NULL, 'Belgium', '1991-05-08', 'f52f876c-832f-4af0-b6e8-7fc93c047a37');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Matt Campbell', NULL, 'Australia', '1995-02-17', '75bf7732-5510-4fb0-8e4d-cb093c04c1bf');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Neel Jani', NULL, 'Switzerland', '1983-12-08', 'e79027fa-44ca-4f3a-b948-6b92931371c2');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Nico Pino', NULL, 'Chile', '2004-09-20', 'b2983bcf-c4bb-4678-8215-9ece58afa433');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Nicolas Varrone', NULL, 'Argentina', '2000-09-25', 'd1e844b7-5a90-46f1-b7d8-e36f0e5509b0');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Will Stevens', NULL, 'United Kingdom', '1991-05-28', '6840fcc9-fec8-4c8d-9cb8-178b1c96ccef');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Alex Lynn', NULL, 'United Kingdom', '1993-09-17', '08f4e629-c708-4b9f-a980-43561d9f1267');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Norman Nato', NULL, 'France', '1992-07-08', 'eccc8ea0-17a6-49e1-83ac-b677503c9806');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Jenson Button', NULL, 'United Kingdom', '1980-01-19', '9694de26-25b9-4a57-adb4-773023bc2740');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Sébastien Bourdais', NULL, 'France', '1979-02-28', 'c11ab2b7-efdb-4a81-b24a-a4d4e561040e');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Earl Bamber', NULL, 'New Zealand', '1990-07-09', '0e82a0b2-8ea0-4569-8b3a-99365e6f44fd');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Dries Vanthoor', NULL, 'Belgium', '1998-04-20', '04a9ad4b-7dea-422c-a9b1-58108285e961');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Raffaele Marciello', NULL, 'Italy', '1994-12-17', 'f011403a-8b67-47c7-89e0-2798dac2d356');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Kevin Magnussen', NULL, 'Denmark', '1992-10-05', '687eef26-c1c1-4e8b-90b9-b20497f81d09');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Sheldon van der Linde', NULL, 'South Africa', '1999-05-13', '8d83f0a6-4b6c-4bc8-8b3e-c4d8a6bffa49');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('René Rast', NULL, 'Germany', '1986-10-26', 'cfa420ab-bf0c-48be-bc8d-c8f473ae5a75');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Robin Frijns', NULL, 'Netherlands', '1991-08-07', 'a4a3225c-d29e-482d-9180-6f17a5943334');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Ferdinand Habsburg', NULL, 'Austria', '1997-06-21', '2e0cf678-3749-4c9c-8363-a4b585f44a44');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Charles Milesi', NULL, 'France', '2001-04-24', '28dd7ee0-f089-4743-acea-50065e5baf00');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Paul-Loup Chatin', NULL, 'France', '1991-10-19', '9e0c8893-5964-45ab-8781-27c0e54f6d1c');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Mick Schumacher', NULL, 'Germany', '1999-03-22', 'd56787ff-9dda-4c5d-a57c-e41567c8cb03');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Jules Gounon', NULL, 'France', '1994-12-31', '69ca8299-d457-4995-b93b-23cad4dd9a2c');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Frédéric Makowiecki', NULL, 'France', '1980-11-22', '14120944-7417-4bce-8073-d97c4c348e2d');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Jean-Éric Vergne', NULL, 'France', '1990-04-25', '08dd3fcf-c817-45c5-b50d-4edd52cc0349');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Mikkel Jensen', NULL, 'Denmark', '1994-12-31', '2ade3470-e6f3-4ac4-8856-4b74e61b495d');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Paul di Resta', NULL, 'United Kingdom', '1986-04-16', '2af51bd9-70e1-414d-b914-6780a8cba253');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Stoffel Vandoorne', NULL, 'Belgium', '1992-03-26', '595977f4-93cc-453e-ae47-78a51143b040');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Loïc Duval', NULL, 'France', '1982-06-12', '2c3cfd0e-4577-4e2a-a51e-e96ab4721bb0');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Malthe Jakobsen', NULL, 'Denmark', '2003-09-04', 'd899e2b6-1adf-4cd1-84ea-7dc3412ef6ef');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Harry Tincknell', NULL, 'United Kingdom', '1991-10-29', 'ec1a153f-8e8b-43c4-8cf3-de0351b14995');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Tom Gamble', NULL, 'United Kingdom', '2001-11-14', 'cda9f220-babf-4002-b7bc-0d00742ba726');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Ross Gunn', NULL, 'United Kingdom', '1997-01-01', '0554e578-f8ec-4746-95cc-a823aafe97ca');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Alex Riberas', NULL, 'Spain', '1994-01-27', '3cc7bd80-3cfe-4247-8e45-a7e3d1f68b89');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Marco Sørensen', NULL, 'Denmark', '1990-09-06', '13a33ba6-d793-4712-aced-0e71831f2399');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Roman De Angelis', NULL, 'Canada', '2001-01-30', '2c1cb73d-3196-41b2-81a0-74053301bef5');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Derek Deboer', NULL, 'United States', '1975-03-04', 'f8379d1e-a511-4629-96ef-99c418bfdf28');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Eduardo Barrichello', NULL, 'Brazil', '2001-09-23', '0a420565-c6e4-4169-b7a5-2cdbe92e43d2');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Valentin Hasse-Clot', NULL, 'France', '2004-06-15', 'd2f1c683-2172-44b7-82a6-3a30d0ed8878');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('François Heriau', NULL, 'France', '1983-10-25', '24477827-d91e-47f5-9706-44659d5c65f2');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Simon Mann', NULL, 'France', '2001-02-01', 'e6e630f6-3df8-4be3-aa51-fd2f97d31826');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Alessio Rovera', NULL, 'Italy', '1995-06-22', '1a19b7fa-ac12-40ce-a014-d78ee255902c');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Ian James', NULL, 'United Kingdom', '1974-07-22', 'c4b6a6ca-aee8-4367-9fdf-afe3fe6cba4e');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Zacharie Robichon', NULL, 'Canada', '1992-06-24', 'a5beded6-7fae-4a55-b6bd-39f72a8d832e');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Mattia Drudi', NULL, 'Italy', '1998-08-10', 'd814cc43-a39a-4e27-be27-61bd506cff97');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Augusto Farfus', NULL, 'Brazil', '1983-09-03', '29abb75e-f4c7-4169-820e-254e42beb140');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Timur Boguslavskiy', NULL, 'Russia', '1998-12-23', 'b309240b-172c-4b66-b3f3-081e8a8fb3b8');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Yasser Shahin', NULL, 'Australia', '1976-09-24', 'a2b726ea-f86c-41fe-b72e-683090016665');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Jonny Edgar', NULL, 'United Kingdom', '2004-02-13', 'c7fecf5e-2cc8-4413-add8-4f12d191f64a');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Daniel Juncadella', NULL, 'Spain', '1991-05-07', '3e1175b3-51e1-4c4f-ab35-86f86feec927');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Ben Keating', NULL, 'United States', '1971-08-18', '617ed046-0e5a-4fb6-b919-4709ee458980');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Kelvin van der Linde', NULL, 'South Africa', '1996-06-20', '5f68a8f7-d016-4d82-bf7a-8fc2e5665ebb');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Valentino Rossi', NULL, 'Italy', '1979-02-16', '2e3cdac7-8754-42aa-8451-5457d413f1bb');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Ahmad Al Harthy', NULL, 'Oman', '1981-08-31', 'ee9b3b5e-9b17-4169-aa55-cc6c9f4f4e52');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Thomas Flohr', NULL, 'Switzerland', '1960-04-17', 'fdf81545-99a3-4942-bdbf-2d20d0b2c9aa');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Francesco Castellacci', NULL, 'Italy', '1987-04-04', 'e10a6e3e-7886-43ac-93ea-b53ca89257cf');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Davide Rigon', NULL, 'Italy', '1986-08-26', 'ad364e55-bab9-4575-becd-21dd849263e9');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('James Cottingham', NULL, 'United Kingdom', '1984-01-29', '17597188-d634-40ab-b901-cbea7095726d');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Sébastien Baud', NULL, 'France', '2000-08-21', '64254f17-7f67-4ea1-8122-b0956c6d5c07');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Grégoire Saucy', NULL, 'Switzerland', '1999-12-26', 'd11f74e9-15b3-40b1-bd13-969b6efd62f4');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Claudio Schiavoni', NULL, 'Italy', '1968-06-27', 'c760613b-c656-41f4-947a-8ec824aeea80');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Matteo Cressoni', NULL, 'Italy', '1984-10-28', '1c8402b4-27fa-4596-b7ab-fca1a0c98020');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Matteo Cairoli', NULL, 'Italy', '1996-06-11', '9f40e026-bff3-4033-875e-fc886125f6dc');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Maxime Martin', NULL, 'Belgium', '1986-03-20', '18ff0bf9-19e3-4765-be9b-409e9f8e4aab');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Christian Ried', NULL, 'Germany', '1979-02-24', '9122eaa2-8945-432c-a9c0-bc22d2a0ee76');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Lin Hodenius', NULL, 'Netherlands', '2006-03-18', 'ea2f3832-1c51-4b67-a30b-465b37a8da59');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Ben Barker', NULL, 'United Kingdom', '1991-04-23', '507028e9-bd69-4d0c-a926-3ef0583327aa');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Bernardo Sousa', NULL, 'Portugal', '1987-03-19', '7908afda-1c81-4c7a-865d-3a6acaf6eada');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Ben Tuck', NULL, 'United Kingdom', '1997-01-24', '7278b7be-3738-4348-bb35-f60eec9425e4');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Finn Gehrsitz', NULL, 'Germany', '2003-06-07', 'a7f771ac-54b2-4d01-9a4b-db05cfa56e80');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Ben Barnicoat', NULL, 'United Kingdom', '1996-12-20', '79ab0dc9-4348-449b-80f7-ccb8299695df');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Arnold Robin', NULL, 'France', '2001-08-21', 'd6b8dd33-d6e7-4322-b2a7-ed0b46ab6623');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Tom van Rompuy', NULL, 'Belgium', '1987-09-08', '8be224ad-13e2-4580-9f73-4e87366883c5');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Rui Andrade', NULL, 'Portugal', '1999-09-23', '5f372921-dbdd-4f92-b630-7930f6e96cf7');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Charlie Eastwood', NULL, 'Ireland', '1995-08-11', '82e5f157-7eb7-44b3-af7a-020825def80c');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Célia Martin', NULL, 'France', '2001-07-14', 'f2496c6f-fff7-4413-ab81-f014766aa871');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Rahel Frey', NULL, 'Switzerland', '1986-02-23', '69a376d0-bf40-4550-ae3a-03b9c3d94a26');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Michelle Gatting', NULL, 'Denmark', '1993-12-06', 'a1c119e0-cc37-42fb-adec-4fdefb4a0c0a');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('José María López', NULL, 'Argentina', '1983-04-26', '697d287b-bfa9-4fdb-a355-d8972901780e');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Clemens Schmid', NULL, 'Austria', '1990-07-08', '7bd90695-ae0a-4318-9817-3bf55eb4138e');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Petru Umbrarescu', NULL, 'Romania', '1993-06-30', '4ac40f48-e466-4bd7-96cd-c12967998f93');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Dennis Olsen', NULL, 'Norway', '1996-04-14', 'f46f108e-bb08-48fc-b504-b01f3d346da8');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Stefano Gattuso', NULL, 'Italy', '1984-05-03', '13ab38a4-a673-4714-a4e8-b2a2798876ab');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Giammarco Levorato', NULL, 'Italy', '2001-05-17', '77ff453d-9312-4696-a1ae-86b79118a440');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Richard Lietz', NULL, 'Austria', '1983-12-17', '694101ca-2a14-4295-a232-aa96999459b3');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Ryan Hardwick', NULL, 'United States', '1980-11-03', '81446b3b-50d6-42c6-931f-e78f3daf2585');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Riccardo Pera', NULL, 'Italy', '1992-08-13', '2a946a2f-4712-4a76-b94d-26c67543ff75');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Marino Sato', NULL, 'Japan', '1999-05-12', 'de470587-b265-47c1-99ac-fc7fe91eb441');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Sean Gelael', NULL, 'Indonesia', '1996-11-01', 'c90a607f-70ff-4113-8682-ad40692db421');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Darren Leung', NULL, 'United Kingdom', '1987-09-25', '255809e5-3d8f-4a2f-a5ac-bb773efec9f4');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Mirko Bortolotti', NULL, 'Italy', '1990-01-10', '6982f216-626b-4c00-8846-5a5cbff8008f');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Daniil Kvyat', NULL, 'Russia', '1994-04-26', 'f8cae580-7786-47a9-a9b6-44fd0e758264');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Edoardo Mortara', NULL, 'Switzerland', '1987-01-12', '54f25a01-a989-4ebe-8a7d-c41315617582');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Carl Bennett', NULL, 'United States', '2004-01-01', '2327ed6b-c15e-42d9-b41e-3c6a07923e10');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Antonio Serravalle', NULL, 'Canada', '2002-09-18', '8d020096-bb3a-433e-89ff-9671173d6ea4');
INSERT INTO driver (name, photo, nationality, birthday, id) VALUES ('Jean-Karl Vernay', NULL, 'France', '1987-10-31', '8d5ff1cb-3059-4420-9b76-2e2f189349ac');


--
-- Data for Name: track; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO track (name, country, `schema`, lap_length, turns, id) VALUES ('Fuji International Speedway', 'Japan', NULL, 4563, 16, '58a8d1fa-20b9-465e-a150-90c768a4bcdd');
INSERT INTO track (name, country, `schema`, lap_length, turns, id) VALUES ('Suzuka International Racing Course', 'Japan', NULL, 5807, 17, 'aee0933e-7a4f-4dc1-b5ad-a3915e3c15d4');
INSERT INTO track (name, country, `schema`, lap_length, turns, id) VALUES ('Circuit de Spa-Francorchamps', 'Belgium', NULL, 7004, 21, 'e879cb90-8401-463e-9706-eac687cdf0b3');
INSERT INTO track (name, country, `schema`, lap_length, turns, id) VALUES ('Circuit de la Sarthe', 'France', NULL, 13629, 38, '4cff5005-9411-4fd8-bfb2-0450519d5538');
INSERT INTO track (name, country, `schema`, lap_length, turns, id) VALUES ('Bahrain International Circuit', 'Bahrain', NULL, 5411, 15, '68f2928a-2512-4e2d-b44d-c5d405659df1');
INSERT INTO track (name, country, `schema`, lap_length, turns, id) VALUES ('Autodromo Jose Carlos Pace (Interlagos)', 'Brazil', NULL, 4309, 15, 'c1e775ec-dc7b-4830-8c1b-e1788007038d');
INSERT INTO track (name, country, `schema`, lap_length, turns, id) VALUES ('Silverstone Circuit', 'United Kingdom', NULL, 5900, 17, '1379c165-a99b-45b0-bdb3-3c557a4f591b');
INSERT INTO track (name, country, `schema`, lap_length, turns, id) VALUES ('Shanghai International Circuit', 'China', NULL, 5451, 16, '577daebf-6b74-4a66-af56-73106d2d10bf');
INSERT INTO track (name, country, `schema`, lap_length, turns, id) VALUES ('Circuit of the Americas', 'USA', NULL, 5513, 20, '3ed435d9-8df4-499c-971d-72382e038b78');
INSERT INTO track (name, country, `schema`, lap_length, turns, id) VALUES ('Autodromo Nazionale Monza', 'Italy', NULL, 5793, 11, '64c1647c-a18f-41f9-b890-fb1eead801a7');
INSERT INTO track (name, country, `schema`, lap_length, turns, id) VALUES ('Autódromo Internacional do Algarve (Portimão)', 'Portugal', NULL, 4684, 15, 'b1f4750f-a51e-4a2c-92a9-e2ad4da4182c');
INSERT INTO track (name, country, `schema`, lap_length, turns, id) VALUES ('Sebring International Raceway', 'USA', NULL, 6019, 17, '505c84c4-c77d-4e92-84c6-fab2fd3ee9ad');
INSERT INTO track (name, country, `schema`, lap_length, turns, id) VALUES ('Nürburgring', 'Germany', NULL, 5148, 16, '1a00242d-5f0b-4993-ae24-56a53835e155');
INSERT INTO track (name, country, `schema`, lap_length, turns, id) VALUES ('Lusail International Circuit', 'Qatar', NULL, 5380, 16, '2570c99d-4ff0-4592-a698-3326aaa13255');
INSERT INTO track (name, country, `schema`, lap_length, turns, id) VALUES ('Autodromo Internazionale Enzo e Dino Ferrari (Imola)', 'Italy', NULL, 4909, 19, '19e7fd47-5896-4f5f-9f7b-224e8447c74e');


--
-- Data for Name: race; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('6 Hours of Fuji', '2025-09-25', 2, 6, '00c27a06-2722-49cd-b233-cde03d3afc7b', '58a8d1fa-20b9-465e-a150-90c768a4bcdd', '6f79a9bf-5c1a-4e1d-a9ab-304dda6d911f');
INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('6 Hours of Imola', '2025-04-20', 2, 6, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '19e7fd47-5896-4f5f-9f7b-224e8447c74e', '6f79a9bf-5c1a-4e1d-a9ab-304dda6d911f');
INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('TotalEnergies 6 Hours of Spa-Francorchamps', '2025-05-10', 2, 6, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', 'e879cb90-8401-463e-9706-eac687cdf0b3', '6f79a9bf-5c1a-4e1d-a9ab-304dda6d911f');
INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('Rolex 6 Hours of São Paulo', '2025-07-13', 2, 6, 'b2754cb0-86d3-43bd-b938-093730498e43', 'c1e775ec-dc7b-4830-8c1b-e1788007038d', '6f79a9bf-5c1a-4e1d-a9ab-304dda6d911f');
INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('Lone Star Le Mans', '2025-09-07', 2, 6, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '3ed435d9-8df4-499c-971d-72382e038b78', '6f79a9bf-5c1a-4e1d-a9ab-304dda6d911f');
INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('6 Hours of Imola', '2024-04-21', 2, 6, '59adf65a-44b8-4fd3-8321-390f32c7b4e1', '19e7fd47-5896-4f5f-9f7b-224e8447c74e', '097b8e91-3e3c-44a4-adba-8838097e05e5');
INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('TotalEnergies 6 Hours of Spa-Francorchamps', '2024-05-11', 2, 6, '1f323ec5-73dc-4843-a772-fe276b0703e3', 'e879cb90-8401-463e-9706-eac687cdf0b3', '097b8e91-3e3c-44a4-adba-8838097e05e5');
INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('6 Hours of São Paulo', '2024-07-14', 2, 6, '8e3b3d1c-6695-4eb5-8bf2-d2f19043a872', 'c1e775ec-dc7b-4830-8c1b-e1788007038d', '097b8e91-3e3c-44a4-adba-8838097e05e5');
INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('Lone Star Le Mans', '2024-09-01', 2, 6, 'b025a79f-90e5-40f4-9387-1694a57844b5', '3ed435d9-8df4-499c-971d-72382e038b78', '097b8e91-3e3c-44a4-adba-8838097e05e5');
INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('6 Hours of Fuji', '2024-09-15', 2, 6, '93ed37ef-7293-4bcd-9531-8a47a79d326b', '58a8d1fa-20b9-465e-a150-90c768a4bcdd', '097b8e91-3e3c-44a4-adba-8838097e05e5');
INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('Qatar 1812Km', '2025-02-28', 2, 10, '50086d13-1588-4502-b636-91fdaa389383', '2570c99d-4ff0-4592-a698-3326aaa13255', '6f79a9bf-5c1a-4e1d-a9ab-304dda6d911f');
INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('Qatar 1812Km', '2024-03-02', 2, 10, '04d32afa-4b87-4272-bf46-abee6c8e48ca', '2570c99d-4ff0-4592-a698-3326aaa13255', '097b8e91-3e3c-44a4-adba-8838097e05e5');
INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('24 Hours of Le Mans', '2025-06-14', 2, 24, '4753da99-54e9-4807-93b5-2e26a4d98d29', '4cff5005-9411-4fd8-bfb2-0450519d5538', '6f79a9bf-5c1a-4e1d-a9ab-304dda6d911f');
INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('24 Hours of Le Mans', '2024-06-15', 2, 24, '0c2a783d-084b-4875-8443-3a5356787c79', '4cff5005-9411-4fd8-bfb2-0450519d5538', '097b8e91-3e3c-44a4-adba-8838097e05e5');
INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('Bapco Energies 8 Hours of Bahrain', '2025-11-08', 2, 8, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '68f2928a-2512-4e2d-b44d-c5d405659df1', '6f79a9bf-5c1a-4e1d-a9ab-304dda6d911f');
INSERT INTO race (name, date, type, duration, id, track, championship) VALUES ('8 Hours of Bahrain', '2024-11-02', 2, 8, '5c272bb3-5b9a-4564-adde-e942d17d3fb3', '68f2928a-2512-4e2d-b44d-c5d405659df1', '097b8e91-3e3c-44a4-adba-8838097e05e5');


--
-- Data for Name: finish; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO finish (pos, race, car_p) VALUES (2, '00c27a06-2722-49cd-b233-cde03d3afc7b', 'd8c0d728-08cb-450a-b5ee-3b98027e80d3');
INSERT INTO finish (pos, race, car_p) VALUES (3, '00c27a06-2722-49cd-b233-cde03d3afc7b', 'c76df0ff-11ca-4170-b88a-1291f8442066');
INSERT INTO finish (pos, race, car_p) VALUES (4, '00c27a06-2722-49cd-b233-cde03d3afc7b', '44fff04f-d19a-4aab-8ee4-a7cdcfc38fd8');
INSERT INTO finish (pos, race, car_p) VALUES (5, '00c27a06-2722-49cd-b233-cde03d3afc7b', '80923a54-81ab-40e7-b9d5-e7da52d83761');
INSERT INTO finish (pos, race, car_p) VALUES (6, '00c27a06-2722-49cd-b233-cde03d3afc7b', 'ddff5320-64f6-460b-a76a-19488ea08e70');
INSERT INTO finish (pos, race, car_p) VALUES (7, '00c27a06-2722-49cd-b233-cde03d3afc7b', 'e03d5b31-1b44-4375-b09c-a303cbbf5d8f');
INSERT INTO finish (pos, race, car_p) VALUES (8, '00c27a06-2722-49cd-b233-cde03d3afc7b', '8793bec9-4be7-459a-b9da-0dd56a223de6');
INSERT INTO finish (pos, race, car_p) VALUES (10, '00c27a06-2722-49cd-b233-cde03d3afc7b', '011f7d02-b781-443c-bccc-c3d61aab0f24');
INSERT INTO finish (pos, race, car_p) VALUES (11, '00c27a06-2722-49cd-b233-cde03d3afc7b', '296e1bc6-c1c7-4261-b48a-935986950265');
INSERT INTO finish (pos, race, car_p) VALUES (13, '00c27a06-2722-49cd-b233-cde03d3afc7b', '147902c8-1f9c-4ed3-a310-91935ee7b41d');
INSERT INTO finish (pos, race, car_p) VALUES (14, '00c27a06-2722-49cd-b233-cde03d3afc7b', '22984b36-d8fd-48f8-8068-006140dd34da');
INSERT INTO finish (pos, race, car_p) VALUES (16, '00c27a06-2722-49cd-b233-cde03d3afc7b', '00b26904-d63c-45aa-a917-1878c7af767e');
INSERT INTO finish (pos, race, car_p) VALUES (17, '00c27a06-2722-49cd-b233-cde03d3afc7b', 'e9da05f6-4335-493e-8b1a-9d050ce33e59');
INSERT INTO finish (pos, race, car_p) VALUES (18, '00c27a06-2722-49cd-b233-cde03d3afc7b', '3f8291ff-6434-400b-bc28-53a15df6bd18');
INSERT INTO finish (pos, race, car_p) VALUES (1, '50086d13-1588-4502-b636-91fdaa389383', '1e1c1666-982b-46f8-9c4a-4350c0994aac');
INSERT INTO finish (pos, race, car_p) VALUES (2, '50086d13-1588-4502-b636-91fdaa389383', 'b1d55df9-d7b0-4f95-bd08-ba8009ba1b29');
INSERT INTO finish (pos, race, car_p) VALUES (3, '50086d13-1588-4502-b636-91fdaa389383', 'f9dc7fe7-bc27-4921-b665-96d1aa72ec87');
INSERT INTO finish (pos, race, car_p) VALUES (4, '50086d13-1588-4502-b636-91fdaa389383', '3f8291ff-6434-400b-bc28-53a15df6bd18');
INSERT INTO finish (pos, race, car_p) VALUES (5, '50086d13-1588-4502-b636-91fdaa389383', '00b26904-d63c-45aa-a917-1878c7af767e');
INSERT INTO finish (pos, race, car_p) VALUES (6, '50086d13-1588-4502-b636-91fdaa389383', 'e03d5b31-1b44-4375-b09c-a303cbbf5d8f');
INSERT INTO finish (pos, race, car_p) VALUES (7, '50086d13-1588-4502-b636-91fdaa389383', '8793bec9-4be7-459a-b9da-0dd56a223de6');
INSERT INTO finish (pos, race, car_p) VALUES (8, '50086d13-1588-4502-b636-91fdaa389383', 'ddff5320-64f6-460b-a76a-19488ea08e70');
INSERT INTO finish (pos, race, car_p) VALUES (9, '50086d13-1588-4502-b636-91fdaa389383', 'd8c0d728-08cb-450a-b5ee-3b98027e80d3');
INSERT INTO finish (pos, race, car_p) VALUES (10, '50086d13-1588-4502-b636-91fdaa389383', '44fff04f-d19a-4aab-8ee4-a7cdcfc38fd8');
INSERT INTO finish (pos, race, car_p) VALUES (11, '50086d13-1588-4502-b636-91fdaa389383', 'c76df0ff-11ca-4170-b88a-1291f8442066');
INSERT INTO finish (pos, race, car_p) VALUES (12, '50086d13-1588-4502-b636-91fdaa389383', '011f7d02-b781-443c-bccc-c3d61aab0f24');
INSERT INTO finish (pos, race, car_p) VALUES (13, '50086d13-1588-4502-b636-91fdaa389383', '22984b36-d8fd-48f8-8068-006140dd34da');
INSERT INTO finish (pos, race, car_p) VALUES (14, '50086d13-1588-4502-b636-91fdaa389383', '60a41976-bb4d-4ce5-ab9c-da8a09106389');
INSERT INTO finish (pos, race, car_p) VALUES (15, '50086d13-1588-4502-b636-91fdaa389383', '296e1bc6-c1c7-4261-b48a-935986950265');
INSERT INTO finish (pos, race, car_p) VALUES (16, '50086d13-1588-4502-b636-91fdaa389383', '147902c8-1f9c-4ed3-a310-91935ee7b41d');
INSERT INTO finish (pos, race, car_p) VALUES (17, '50086d13-1588-4502-b636-91fdaa389383', '80923a54-81ab-40e7-b9d5-e7da52d83761');
INSERT INTO finish (pos, race, car_p) VALUES (18, '50086d13-1588-4502-b636-91fdaa389383', 'e9da05f6-4335-493e-8b1a-9d050ce33e59');
INSERT INTO finish (pos, race, car_p) VALUES (1, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', 'f9dc7fe7-bc27-4921-b665-96d1aa72ec87');
INSERT INTO finish (pos, race, car_p) VALUES (2, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '8793bec9-4be7-459a-b9da-0dd56a223de6');
INSERT INTO finish (pos, race, car_p) VALUES (3, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '22984b36-d8fd-48f8-8068-006140dd34da');
INSERT INTO finish (pos, race, car_p) VALUES (4, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', 'b1d55df9-d7b0-4f95-bd08-ba8009ba1b29');
INSERT INTO finish (pos, race, car_p) VALUES (5, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '00b26904-d63c-45aa-a917-1878c7af767e');
INSERT INTO finish (pos, race, car_p) VALUES (6, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '3f8291ff-6434-400b-bc28-53a15df6bd18');
INSERT INTO finish (pos, race, car_p) VALUES (7, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', 'e03d5b31-1b44-4375-b09c-a303cbbf5d8f');
INSERT INTO finish (pos, race, car_p) VALUES (8, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', 'c76df0ff-11ca-4170-b88a-1291f8442066');
INSERT INTO finish (pos, race, car_p) VALUES (9, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', 'd8c0d728-08cb-450a-b5ee-3b98027e80d3');
INSERT INTO finish (pos, race, car_p) VALUES (10, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', 'ddff5320-64f6-460b-a76a-19488ea08e70');
INSERT INTO finish (pos, race, car_p) VALUES (11, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '44fff04f-d19a-4aab-8ee4-a7cdcfc38fd8');
INSERT INTO finish (pos, race, car_p) VALUES (12, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '011f7d02-b781-443c-bccc-c3d61aab0f24');
INSERT INTO finish (pos, race, car_p) VALUES (13, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '60a41976-bb4d-4ce5-ab9c-da8a09106389');
INSERT INTO finish (pos, race, car_p) VALUES (14, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '296e1bc6-c1c7-4261-b48a-935986950265');
INSERT INTO finish (pos, race, car_p) VALUES (15, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '1e1c1666-982b-46f8-9c4a-4350c0994aac');
INSERT INTO finish (pos, race, car_p) VALUES (16, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '147902c8-1f9c-4ed3-a310-91935ee7b41d');
INSERT INTO finish (pos, race, car_p) VALUES (17, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '80923a54-81ab-40e7-b9d5-e7da52d83761');
INSERT INTO finish (pos, race, car_p) VALUES (18, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', 'e9da05f6-4335-493e-8b1a-9d050ce33e59');
INSERT INTO finish (pos, race, car_p) VALUES (1, '4753da99-54e9-4807-93b5-2e26a4d98d29', 'b1d55df9-d7b0-4f95-bd08-ba8009ba1b29');
INSERT INTO finish (pos, race, car_p) VALUES (2, '4753da99-54e9-4807-93b5-2e26a4d98d29', 'c76df0ff-11ca-4170-b88a-1291f8442066');
INSERT INTO finish (pos, race, car_p) VALUES (3, '4753da99-54e9-4807-93b5-2e26a4d98d29', 'f9dc7fe7-bc27-4921-b665-96d1aa72ec87');
INSERT INTO finish (pos, race, car_p) VALUES (4, '4753da99-54e9-4807-93b5-2e26a4d98d29', 'ddff5320-64f6-460b-a76a-19488ea08e70');
INSERT INTO finish (pos, race, car_p) VALUES (5, '4753da99-54e9-4807-93b5-2e26a4d98d29', 'e03d5b31-1b44-4375-b09c-a303cbbf5d8f');
INSERT INTO finish (pos, race, car_p) VALUES (6, '4753da99-54e9-4807-93b5-2e26a4d98d29', '44fff04f-d19a-4aab-8ee4-a7cdcfc38fd8');
INSERT INTO finish (pos, race, car_p) VALUES (7, '4753da99-54e9-4807-93b5-2e26a4d98d29', '147902c8-1f9c-4ed3-a310-91935ee7b41d');
INSERT INTO finish (pos, race, car_p) VALUES (9, '4753da99-54e9-4807-93b5-2e26a4d98d29', '60a41976-bb4d-4ce5-ab9c-da8a09106389');
INSERT INTO finish (pos, race, car_p) VALUES (10, '4753da99-54e9-4807-93b5-2e26a4d98d29', '22984b36-d8fd-48f8-8068-006140dd34da');
INSERT INTO finish (pos, race, car_p) VALUES (11, '4753da99-54e9-4807-93b5-2e26a4d98d29', '011f7d02-b781-443c-bccc-c3d61aab0f24');
INSERT INTO finish (pos, race, car_p) VALUES (12, '4753da99-54e9-4807-93b5-2e26a4d98d29', '80923a54-81ab-40e7-b9d5-e7da52d83761');
INSERT INTO finish (pos, race, car_p) VALUES (13, '4753da99-54e9-4807-93b5-2e26a4d98d29', '296e1bc6-c1c7-4261-b48a-935986950265');
INSERT INTO finish (pos, race, car_p) VALUES (14, '4753da99-54e9-4807-93b5-2e26a4d98d29', 'e9da05f6-4335-493e-8b1a-9d050ce33e59');
INSERT INTO finish (pos, race, car_p) VALUES (15, '4753da99-54e9-4807-93b5-2e26a4d98d29', '00b26904-d63c-45aa-a917-1878c7af767e');
INSERT INTO finish (pos, race, car_p) VALUES (16, '4753da99-54e9-4807-93b5-2e26a4d98d29', 'd8c0d728-08cb-450a-b5ee-3b98027e80d3');
INSERT INTO finish (pos, race, car_p) VALUES (17, '4753da99-54e9-4807-93b5-2e26a4d98d29', '8793bec9-4be7-459a-b9da-0dd56a223de6');
INSERT INTO finish (pos, race, car_p) VALUES (18, '4753da99-54e9-4807-93b5-2e26a4d98d29', '3f8291ff-6434-400b-bc28-53a15df6bd18');
INSERT INTO finish (pos, race, car_p) VALUES (21, '4753da99-54e9-4807-93b5-2e26a4d98d29', '1e1c1666-982b-46f8-9c4a-4350c0994aac');
INSERT INTO finish (pos, race, car_p) VALUES (1, 'b2754cb0-86d3-43bd-b938-093730498e43', 'ddff5320-64f6-460b-a76a-19488ea08e70');
INSERT INTO finish (pos, race, car_p) VALUES (2, 'b2754cb0-86d3-43bd-b938-093730498e43', '147902c8-1f9c-4ed3-a310-91935ee7b41d');
INSERT INTO finish (pos, race, car_p) VALUES (3, 'b2754cb0-86d3-43bd-b938-093730498e43', '44fff04f-d19a-4aab-8ee4-a7cdcfc38fd8');
INSERT INTO finish (pos, race, car_p) VALUES (4, 'b2754cb0-86d3-43bd-b938-093730498e43', 'c76df0ff-11ca-4170-b88a-1291f8442066');
INSERT INTO finish (pos, race, car_p) VALUES (5, 'b2754cb0-86d3-43bd-b938-093730498e43', '8793bec9-4be7-459a-b9da-0dd56a223de6');
INSERT INTO finish (pos, race, car_p) VALUES (6, 'b2754cb0-86d3-43bd-b938-093730498e43', '011f7d02-b781-443c-bccc-c3d61aab0f24');
INSERT INTO finish (pos, race, car_p) VALUES (7, 'b2754cb0-86d3-43bd-b938-093730498e43', 'd8c0d728-08cb-450a-b5ee-3b98027e80d3');
INSERT INTO finish (pos, race, car_p) VALUES (8, 'b2754cb0-86d3-43bd-b938-093730498e43', 'b1d55df9-d7b0-4f95-bd08-ba8009ba1b29');
INSERT INTO finish (pos, race, car_p) VALUES (9, 'b2754cb0-86d3-43bd-b938-093730498e43', '22984b36-d8fd-48f8-8068-006140dd34da');
INSERT INTO finish (pos, race, car_p) VALUES (10, 'b2754cb0-86d3-43bd-b938-093730498e43', '296e1bc6-c1c7-4261-b48a-935986950265');
INSERT INTO finish (pos, race, car_p) VALUES (11, 'b2754cb0-86d3-43bd-b938-093730498e43', 'f9dc7fe7-bc27-4921-b665-96d1aa72ec87');
INSERT INTO finish (pos, race, car_p) VALUES (12, 'b2754cb0-86d3-43bd-b938-093730498e43', '1e1c1666-982b-46f8-9c4a-4350c0994aac');
INSERT INTO finish (pos, race, car_p) VALUES (13, 'b2754cb0-86d3-43bd-b938-093730498e43', '80923a54-81ab-40e7-b9d5-e7da52d83761');
INSERT INTO finish (pos, race, car_p) VALUES (14, 'b2754cb0-86d3-43bd-b938-093730498e43', 'e03d5b31-1b44-4375-b09c-a303cbbf5d8f');
INSERT INTO finish (pos, race, car_p) VALUES (15, 'b2754cb0-86d3-43bd-b938-093730498e43', '00b26904-d63c-45aa-a917-1878c7af767e');
INSERT INTO finish (pos, race, car_p) VALUES (16, 'b2754cb0-86d3-43bd-b938-093730498e43', 'e9da05f6-4335-493e-8b1a-9d050ce33e59');
INSERT INTO finish (pos, race, car_p) VALUES (14, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '9a46a97f-30ad-43a0-aeda-91a280947848');
INSERT INTO finish (pos, race, car_p) VALUES (15, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '28d21e2c-ef12-41e4-ab27-3c17b22c60c1');
INSERT INTO finish (pos, race, car_p) VALUES (16, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', 'd262ce67-5829-4eb6-befa-48d07e08b6c6');
INSERT INTO finish (pos, race, car_p) VALUES (17, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '624b63a0-41d9-4c3b-af56-d631afebf840');
INSERT INTO finish (pos, race, car_p) VALUES (18, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '1323c6db-0a94-4b45-bd07-5a7091e777e9');
INSERT INTO finish (pos, race, car_p) VALUES (1, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', 'f9dc7fe7-bc27-4921-b665-96d1aa72ec87');
INSERT INTO finish (pos, race, car_p) VALUES (2, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '1e1c1666-982b-46f8-9c4a-4350c0994aac');
INSERT INTO finish (pos, race, car_p) VALUES (3, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '22984b36-d8fd-48f8-8068-006140dd34da');
INSERT INTO finish (pos, race, car_p) VALUES (4, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '00b26904-d63c-45aa-a917-1878c7af767e');
INSERT INTO finish (pos, race, car_p) VALUES (5, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', 'ddff5320-64f6-460b-a76a-19488ea08e70');
INSERT INTO finish (pos, race, car_p) VALUES (6, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '147902c8-1f9c-4ed3-a310-91935ee7b41d');
INSERT INTO finish (pos, race, car_p) VALUES (7, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', 'e03d5b31-1b44-4375-b09c-a303cbbf5d8f');
INSERT INTO finish (pos, race, car_p) VALUES (8, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '60a41976-bb4d-4ce5-ab9c-da8a09106389');
INSERT INTO finish (pos, race, car_p) VALUES (9, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', 'c76df0ff-11ca-4170-b88a-1291f8442066');
INSERT INTO finish (pos, race, car_p) VALUES (10, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '3f8291ff-6434-400b-bc28-53a15df6bd18');
INSERT INTO finish (pos, race, car_p) VALUES (11, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', 'd8c0d728-08cb-450a-b5ee-3b98027e80d3');
INSERT INTO finish (pos, race, car_p) VALUES (12, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '44fff04f-d19a-4aab-8ee4-a7cdcfc38fd8');
INSERT INTO finish (pos, race, car_p) VALUES (13, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', 'e9da05f6-4335-493e-8b1a-9d050ce33e59');
INSERT INTO finish (pos, race, car_p) VALUES (14, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '80923a54-81ab-40e7-b9d5-e7da52d83761');
INSERT INTO finish (pos, race, car_p) VALUES (15, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', 'b1d55df9-d7b0-4f95-bd08-ba8009ba1b29');
INSERT INTO finish (pos, race, car_p) VALUES (16, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '8793bec9-4be7-459a-b9da-0dd56a223de6');
INSERT INTO finish (pos, race, car_p) VALUES (17, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '011f7d02-b781-443c-bccc-c3d61aab0f24');
INSERT INTO finish (pos, race, car_p) VALUES (18, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '296e1bc6-c1c7-4261-b48a-935986950265');
INSERT INTO finish (pos, race, car_p) VALUES (17, 'b2754cb0-86d3-43bd-b938-093730498e43', '3f8291ff-6434-400b-bc28-53a15df6bd18');
INSERT INTO finish (pos, race, car_p) VALUES (18, 'b2754cb0-86d3-43bd-b938-093730498e43', '60a41976-bb4d-4ce5-ab9c-da8a09106389');
INSERT INTO finish (pos, race, car_p) VALUES (1, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', 'c76df0ff-11ca-4170-b88a-1291f8442066');
INSERT INTO finish (pos, race, car_p) VALUES (2, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '1e1c1666-982b-46f8-9c4a-4350c0994aac');
INSERT INTO finish (pos, race, car_p) VALUES (3, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '011f7d02-b781-443c-bccc-c3d61aab0f24');
INSERT INTO finish (pos, race, car_p) VALUES (4, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', 'd8c0d728-08cb-450a-b5ee-3b98027e80d3');
INSERT INTO finish (pos, race, car_p) VALUES (5, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', 'f9dc7fe7-bc27-4921-b665-96d1aa72ec87');
INSERT INTO finish (pos, race, car_p) VALUES (6, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '147902c8-1f9c-4ed3-a310-91935ee7b41d');
INSERT INTO finish (pos, race, car_p) VALUES (7, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', 'b1d55df9-d7b0-4f95-bd08-ba8009ba1b29');
INSERT INTO finish (pos, race, car_p) VALUES (8, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', 'ddff5320-64f6-460b-a76a-19488ea08e70');
INSERT INTO finish (pos, race, car_p) VALUES (9, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '00b26904-d63c-45aa-a917-1878c7af767e');
INSERT INTO finish (pos, race, car_p) VALUES (10, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '44fff04f-d19a-4aab-8ee4-a7cdcfc38fd8');
INSERT INTO finish (pos, race, car_p) VALUES (11, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '60a41976-bb4d-4ce5-ab9c-da8a09106389');
INSERT INTO finish (pos, race, car_p) VALUES (12, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '3f8291ff-6434-400b-bc28-53a15df6bd18');
INSERT INTO finish (pos, race, car_p) VALUES (13, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '296e1bc6-c1c7-4261-b48a-935986950265');
INSERT INTO finish (pos, race, car_p) VALUES (14, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', 'e03d5b31-1b44-4375-b09c-a303cbbf5d8f');
INSERT INTO finish (pos, race, car_p) VALUES (15, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '22984b36-d8fd-48f8-8068-006140dd34da');
INSERT INTO finish (pos, race, car_p) VALUES (16, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '8793bec9-4be7-459a-b9da-0dd56a223de6');
INSERT INTO finish (pos, race, car_p) VALUES (17, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '80923a54-81ab-40e7-b9d5-e7da52d83761');
INSERT INTO finish (pos, race, car_p) VALUES (18, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', 'e9da05f6-4335-493e-8b1a-9d050ce33e59');
INSERT INTO finish (pos, race, car_p) VALUES (1, '6b317387-403f-4ad3-9bbc-23cd22fb1310', 'e03d5b31-1b44-4375-b09c-a303cbbf5d8f');
INSERT INTO finish (pos, race, car_p) VALUES (2, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '00b26904-d63c-45aa-a917-1878c7af767e');
INSERT INTO finish (pos, race, car_p) VALUES (3, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '1e1c1666-982b-46f8-9c4a-4350c0994aac');
INSERT INTO finish (pos, race, car_p) VALUES (4, '6b317387-403f-4ad3-9bbc-23cd22fb1310', 'f9dc7fe7-bc27-4921-b665-96d1aa72ec87');
INSERT INTO finish (pos, race, car_p) VALUES (5, '6b317387-403f-4ad3-9bbc-23cd22fb1310', 'b1d55df9-d7b0-4f95-bd08-ba8009ba1b29');
INSERT INTO finish (pos, race, car_p) VALUES (6, '6b317387-403f-4ad3-9bbc-23cd22fb1310', 'ddff5320-64f6-460b-a76a-19488ea08e70');
INSERT INTO finish (pos, race, car_p) VALUES (7, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '80923a54-81ab-40e7-b9d5-e7da52d83761');
INSERT INTO finish (pos, race, car_p) VALUES (8, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '8793bec9-4be7-459a-b9da-0dd56a223de6');
INSERT INTO finish (pos, race, car_p) VALUES (9, '6b317387-403f-4ad3-9bbc-23cd22fb1310', 'd8c0d728-08cb-450a-b5ee-3b98027e80d3');
INSERT INTO finish (pos, race, car_p) VALUES (10, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '011f7d02-b781-443c-bccc-c3d61aab0f24');
INSERT INTO finish (pos, race, car_p) VALUES (11, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '60a41976-bb4d-4ce5-ab9c-da8a09106389');
INSERT INTO finish (pos, race, car_p) VALUES (12, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '22984b36-d8fd-48f8-8068-006140dd34da');
INSERT INTO finish (pos, race, car_p) VALUES (13, '6b317387-403f-4ad3-9bbc-23cd22fb1310', 'c76df0ff-11ca-4170-b88a-1291f8442066');
INSERT INTO finish (pos, race, car_p) VALUES (14, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '44fff04f-d19a-4aab-8ee4-a7cdcfc38fd8');
INSERT INTO finish (pos, race, car_p) VALUES (15, '6b317387-403f-4ad3-9bbc-23cd22fb1310', 'e9da05f6-4335-493e-8b1a-9d050ce33e59');
INSERT INTO finish (pos, race, car_p) VALUES (16, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '147902c8-1f9c-4ed3-a310-91935ee7b41d');
INSERT INTO finish (pos, race, car_p) VALUES (17, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '296e1bc6-c1c7-4261-b48a-935986950265');
INSERT INTO finish (pos, race, car_p) VALUES (18, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '3f8291ff-6434-400b-bc28-53a15df6bd18');
INSERT INTO finish (pos, race, car_p) VALUES (1, '50086d13-1588-4502-b636-91fdaa389383', 'd2335551-2808-47bb-8229-6cd849c3af1a');
INSERT INTO finish (pos, race, car_p) VALUES (2, '50086d13-1588-4502-b636-91fdaa389383', '9a46a97f-30ad-43a0-aeda-91a280947848');
INSERT INTO finish (pos, race, car_p) VALUES (3, '50086d13-1588-4502-b636-91fdaa389383', 'c65c3f38-77fd-4099-a10d-35d75d6be870');
INSERT INTO finish (pos, race, car_p) VALUES (4, '50086d13-1588-4502-b636-91fdaa389383', '70645b62-8df4-4334-a8d5-79acce411cba');
INSERT INTO finish (pos, race, car_p) VALUES (5, '50086d13-1588-4502-b636-91fdaa389383', '624b63a0-41d9-4c3b-af56-d631afebf840');
INSERT INTO finish (pos, race, car_p) VALUES (6, '50086d13-1588-4502-b636-91fdaa389383', '1323c6db-0a94-4b45-bd07-5a7091e777e9');
INSERT INTO finish (pos, race, car_p) VALUES (7, '50086d13-1588-4502-b636-91fdaa389383', '1461d107-ea45-4dbe-b456-92cba175a824');
INSERT INTO finish (pos, race, car_p) VALUES (8, '50086d13-1588-4502-b636-91fdaa389383', '1655d35e-8a04-47eb-b694-fcc0f8b8c968');
INSERT INTO finish (pos, race, car_p) VALUES (9, '50086d13-1588-4502-b636-91fdaa389383', '92116a61-59c8-4b83-ab9e-f4153f3c9337');
INSERT INTO finish (pos, race, car_p) VALUES (10, '50086d13-1588-4502-b636-91fdaa389383', 'd262ce67-5829-4eb6-befa-48d07e08b6c6');
INSERT INTO finish (pos, race, car_p) VALUES (11, '50086d13-1588-4502-b636-91fdaa389383', '621245bb-31ee-4da9-b224-6e30be870e17');
INSERT INTO finish (pos, race, car_p) VALUES (12, '50086d13-1588-4502-b636-91fdaa389383', '716d317e-bde0-4fa3-9738-ec0ef9c3cce1');
INSERT INTO finish (pos, race, car_p) VALUES (13, '50086d13-1588-4502-b636-91fdaa389383', 'f334e8a1-28fd-416a-a0ce-dad0423e385b');
INSERT INTO finish (pos, race, car_p) VALUES (14, '50086d13-1588-4502-b636-91fdaa389383', '28d21e2c-ef12-41e4-ab27-3c17b22c60c1');
INSERT INTO finish (pos, race, car_p) VALUES (15, '50086d13-1588-4502-b636-91fdaa389383', '9a97bead-1061-4cc7-85a2-866670c59342');
INSERT INTO finish (pos, race, car_p) VALUES (16, '50086d13-1588-4502-b636-91fdaa389383', 'f15657ca-919a-420d-8785-f605ce6098ab');
INSERT INTO finish (pos, race, car_p) VALUES (17, '50086d13-1588-4502-b636-91fdaa389383', 'c2afd8f8-285f-4903-982a-770a2eacfa04');
INSERT INTO finish (pos, race, car_p) VALUES (18, '50086d13-1588-4502-b636-91fdaa389383', 'b5f023e6-d2c9-4e10-841c-536f5eb176eb');
INSERT INTO finish (pos, race, car_p) VALUES (1, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '716d317e-bde0-4fa3-9738-ec0ef9c3cce1');
INSERT INTO finish (pos, race, car_p) VALUES (2, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '621245bb-31ee-4da9-b224-6e30be870e17');
INSERT INTO finish (pos, race, car_p) VALUES (3, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '70645b62-8df4-4334-a8d5-79acce411cba');
INSERT INTO finish (pos, race, car_p) VALUES (6, '4753da99-54e9-4807-93b5-2e26a4d98d29', 'd2335551-2808-47bb-8229-6cd849c3af1a');
INSERT INTO finish (pos, race, car_p) VALUES (7, '4753da99-54e9-4807-93b5-2e26a4d98d29', '9a97bead-1061-4cc7-85a2-866670c59342');
INSERT INTO finish (pos, race, car_p) VALUES (1, '4753da99-54e9-4807-93b5-2e26a4d98d29', '716d317e-bde0-4fa3-9738-ec0ef9c3cce1');
INSERT INTO finish (pos, race, car_p) VALUES (2, '4753da99-54e9-4807-93b5-2e26a4d98d29', '624b63a0-41d9-4c3b-af56-d631afebf840');
INSERT INTO finish (pos, race, car_p) VALUES (3, '4753da99-54e9-4807-93b5-2e26a4d98d29', 'c2afd8f8-285f-4903-982a-770a2eacfa04');
INSERT INTO finish (pos, race, car_p) VALUES (4, '4753da99-54e9-4807-93b5-2e26a4d98d29', '1323c6db-0a94-4b45-bd07-5a7091e777e9');
INSERT INTO finish (pos, race, car_p) VALUES (5, '4753da99-54e9-4807-93b5-2e26a4d98d29', 'b5f023e6-d2c9-4e10-841c-536f5eb176eb');
INSERT INTO finish (pos, race, car_p) VALUES (8, '4753da99-54e9-4807-93b5-2e26a4d98d29', 'f15657ca-919a-420d-8785-f605ce6098ab');
INSERT INTO finish (pos, race, car_p) VALUES (9, '4753da99-54e9-4807-93b5-2e26a4d98d29', '92116a61-59c8-4b83-ab9e-f4153f3c9337');
INSERT INTO finish (pos, race, car_p) VALUES (10, '4753da99-54e9-4807-93b5-2e26a4d98d29', 'f334e8a1-28fd-416a-a0ce-dad0423e385b');
INSERT INTO finish (pos, race, car_p) VALUES (11, '4753da99-54e9-4807-93b5-2e26a4d98d29', '9a46a97f-30ad-43a0-aeda-91a280947848');
INSERT INTO finish (pos, race, car_p) VALUES (12, '4753da99-54e9-4807-93b5-2e26a4d98d29', '70645b62-8df4-4334-a8d5-79acce411cba');
INSERT INTO finish (pos, race, car_p) VALUES (13, '4753da99-54e9-4807-93b5-2e26a4d98d29', '1655d35e-8a04-47eb-b694-fcc0f8b8c968');
INSERT INTO finish (pos, race, car_p) VALUES (14, '4753da99-54e9-4807-93b5-2e26a4d98d29', 'c65c3f38-77fd-4099-a10d-35d75d6be870');
INSERT INTO finish (pos, race, car_p) VALUES (15, '4753da99-54e9-4807-93b5-2e26a4d98d29', '621245bb-31ee-4da9-b224-6e30be870e17');
INSERT INTO finish (pos, race, car_p) VALUES (16, '4753da99-54e9-4807-93b5-2e26a4d98d29', '1461d107-ea45-4dbe-b456-92cba175a824');
INSERT INTO finish (pos, race, car_p) VALUES (17, '4753da99-54e9-4807-93b5-2e26a4d98d29', '28d21e2c-ef12-41e4-ab27-3c17b22c60c1');
INSERT INTO finish (pos, race, car_p) VALUES (18, '4753da99-54e9-4807-93b5-2e26a4d98d29', 'd262ce67-5829-4eb6-befa-48d07e08b6c6');
INSERT INTO finish (pos, race, car_p) VALUES (9, '00c27a06-2722-49cd-b233-cde03d3afc7b', 'b1d55df9-d7b0-4f95-bd08-ba8009ba1b29');
INSERT INTO finish (pos, race, car_p) VALUES (12, '00c27a06-2722-49cd-b233-cde03d3afc7b', '1e1c1666-982b-46f8-9c4a-4350c0994aac');
INSERT INTO finish (pos, race, car_p) VALUES (15, '00c27a06-2722-49cd-b233-cde03d3afc7b', 'f9dc7fe7-bc27-4921-b665-96d1aa72ec87');
INSERT INTO finish (pos, race, car_p) VALUES (1, '00c27a06-2722-49cd-b233-cde03d3afc7b', '60a41976-bb4d-4ce5-ab9c-da8a09106389');
INSERT INTO finish (pos, race, car_p) VALUES (4, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', 'b5f023e6-d2c9-4e10-841c-536f5eb176eb');
INSERT INTO finish (pos, race, car_p) VALUES (5, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '1655d35e-8a04-47eb-b694-fcc0f8b8c968');
INSERT INTO finish (pos, race, car_p) VALUES (6, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', 'c2afd8f8-285f-4903-982a-770a2eacfa04');
INSERT INTO finish (pos, race, car_p) VALUES (7, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', 'd2335551-2808-47bb-8229-6cd849c3af1a');
INSERT INTO finish (pos, race, car_p) VALUES (8, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', 'f334e8a1-28fd-416a-a0ce-dad0423e385b');
INSERT INTO finish (pos, race, car_p) VALUES (9, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '1461d107-ea45-4dbe-b456-92cba175a824');
INSERT INTO finish (pos, race, car_p) VALUES (10, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '9a97bead-1061-4cc7-85a2-866670c59342');
INSERT INTO finish (pos, race, car_p) VALUES (11, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', '92116a61-59c8-4b83-ab9e-f4153f3c9337');
INSERT INTO finish (pos, race, car_p) VALUES (12, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', 'c65c3f38-77fd-4099-a10d-35d75d6be870');
INSERT INTO finish (pos, race, car_p) VALUES (13, 'e4da4234-4cb0-4cb3-9644-44541c8f0e2a', 'f15657ca-919a-420d-8785-f605ce6098ab');
INSERT INTO finish (pos, race, car_p) VALUES (1, 'b2754cb0-86d3-43bd-b938-093730498e43', 'b5f023e6-d2c9-4e10-841c-536f5eb176eb');
INSERT INTO finish (pos, race, car_p) VALUES (2, 'b2754cb0-86d3-43bd-b938-093730498e43', 'c2afd8f8-285f-4903-982a-770a2eacfa04');
INSERT INTO finish (pos, race, car_p) VALUES (3, 'b2754cb0-86d3-43bd-b938-093730498e43', '92116a61-59c8-4b83-ab9e-f4153f3c9337');
INSERT INTO finish (pos, race, car_p) VALUES (4, 'b2754cb0-86d3-43bd-b938-093730498e43', 'f334e8a1-28fd-416a-a0ce-dad0423e385b');
INSERT INTO finish (pos, race, car_p) VALUES (5, 'b2754cb0-86d3-43bd-b938-093730498e43', '70645b62-8df4-4334-a8d5-79acce411cba');
INSERT INTO finish (pos, race, car_p) VALUES (6, 'b2754cb0-86d3-43bd-b938-093730498e43', '716d317e-bde0-4fa3-9738-ec0ef9c3cce1');
INSERT INTO finish (pos, race, car_p) VALUES (7, 'b2754cb0-86d3-43bd-b938-093730498e43', 'd2335551-2808-47bb-8229-6cd849c3af1a');
INSERT INTO finish (pos, race, car_p) VALUES (8, 'b2754cb0-86d3-43bd-b938-093730498e43', '9a46a97f-30ad-43a0-aeda-91a280947848');
INSERT INTO finish (pos, race, car_p) VALUES (9, 'b2754cb0-86d3-43bd-b938-093730498e43', '1461d107-ea45-4dbe-b456-92cba175a824');
INSERT INTO finish (pos, race, car_p) VALUES (10, 'b2754cb0-86d3-43bd-b938-093730498e43', '621245bb-31ee-4da9-b224-6e30be870e17');
INSERT INTO finish (pos, race, car_p) VALUES (11, 'b2754cb0-86d3-43bd-b938-093730498e43', '1655d35e-8a04-47eb-b694-fcc0f8b8c968');
INSERT INTO finish (pos, race, car_p) VALUES (12, 'b2754cb0-86d3-43bd-b938-093730498e43', 'c65c3f38-77fd-4099-a10d-35d75d6be870');
INSERT INTO finish (pos, race, car_p) VALUES (13, 'b2754cb0-86d3-43bd-b938-093730498e43', '624b63a0-41d9-4c3b-af56-d631afebf840');
INSERT INTO finish (pos, race, car_p) VALUES (14, 'b2754cb0-86d3-43bd-b938-093730498e43', '1323c6db-0a94-4b45-bd07-5a7091e777e9');
INSERT INTO finish (pos, race, car_p) VALUES (15, 'b2754cb0-86d3-43bd-b938-093730498e43', '28d21e2c-ef12-41e4-ab27-3c17b22c60c1');
INSERT INTO finish (pos, race, car_p) VALUES (16, 'b2754cb0-86d3-43bd-b938-093730498e43', 'f15657ca-919a-420d-8785-f605ce6098ab');
INSERT INTO finish (pos, race, car_p) VALUES (17, 'b2754cb0-86d3-43bd-b938-093730498e43', '9a97bead-1061-4cc7-85a2-866670c59342');
INSERT INTO finish (pos, race, car_p) VALUES (18, 'b2754cb0-86d3-43bd-b938-093730498e43', 'd262ce67-5829-4eb6-befa-48d07e08b6c6');
INSERT INTO finish (pos, race, car_p) VALUES (1, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '1461d107-ea45-4dbe-b456-92cba175a824');
INSERT INTO finish (pos, race, car_p) VALUES (2, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '621245bb-31ee-4da9-b224-6e30be870e17');
INSERT INTO finish (pos, race, car_p) VALUES (3, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '1655d35e-8a04-47eb-b694-fcc0f8b8c968');
INSERT INTO finish (pos, race, car_p) VALUES (4, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '9a46a97f-30ad-43a0-aeda-91a280947848');
INSERT INTO finish (pos, race, car_p) VALUES (5, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '1323c6db-0a94-4b45-bd07-5a7091e777e9');
INSERT INTO finish (pos, race, car_p) VALUES (6, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '9a97bead-1061-4cc7-85a2-866670c59342');
INSERT INTO finish (pos, race, car_p) VALUES (7, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '716d317e-bde0-4fa3-9738-ec0ef9c3cce1');
INSERT INTO finish (pos, race, car_p) VALUES (8, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', 'd262ce67-5829-4eb6-befa-48d07e08b6c6');
INSERT INTO finish (pos, race, car_p) VALUES (9, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', 'c65c3f38-77fd-4099-a10d-35d75d6be870');
INSERT INTO finish (pos, race, car_p) VALUES (10, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '28d21e2c-ef12-41e4-ab27-3c17b22c60c1');
INSERT INTO finish (pos, race, car_p) VALUES (11, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '92116a61-59c8-4b83-ab9e-f4153f3c9337');
INSERT INTO finish (pos, race, car_p) VALUES (12, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '624b63a0-41d9-4c3b-af56-d631afebf840');
INSERT INTO finish (pos, race, car_p) VALUES (13, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', 'c2afd8f8-285f-4903-982a-770a2eacfa04');
INSERT INTO finish (pos, race, car_p) VALUES (14, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', '70645b62-8df4-4334-a8d5-79acce411cba');
INSERT INTO finish (pos, race, car_p) VALUES (15, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', 'f15657ca-919a-420d-8785-f605ce6098ab');
INSERT INTO finish (pos, race, car_p) VALUES (16, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', 'd2335551-2808-47bb-8229-6cd849c3af1a');
INSERT INTO finish (pos, race, car_p) VALUES (17, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', 'b5f023e6-d2c9-4e10-841c-536f5eb176eb');
INSERT INTO finish (pos, race, car_p) VALUES (18, '6e4e27f2-a059-4e43-a8ea-a8031faf394b', 'f334e8a1-28fd-416a-a0ce-dad0423e385b');
INSERT INTO finish (pos, race, car_p) VALUES (1, '00c27a06-2722-49cd-b233-cde03d3afc7b', 'c2afd8f8-285f-4903-982a-770a2eacfa04');
INSERT INTO finish (pos, race, car_p) VALUES (2, '00c27a06-2722-49cd-b233-cde03d3afc7b', '624b63a0-41d9-4c3b-af56-d631afebf840');
INSERT INTO finish (pos, race, car_p) VALUES (3, '00c27a06-2722-49cd-b233-cde03d3afc7b', 'c65c3f38-77fd-4099-a10d-35d75d6be870');
INSERT INTO finish (pos, race, car_p) VALUES (4, '00c27a06-2722-49cd-b233-cde03d3afc7b', '621245bb-31ee-4da9-b224-6e30be870e17');
INSERT INTO finish (pos, race, car_p) VALUES (5, '00c27a06-2722-49cd-b233-cde03d3afc7b', '716d317e-bde0-4fa3-9738-ec0ef9c3cce1');
INSERT INTO finish (pos, race, car_p) VALUES (6, '00c27a06-2722-49cd-b233-cde03d3afc7b', '1655d35e-8a04-47eb-b694-fcc0f8b8c968');
INSERT INTO finish (pos, race, car_p) VALUES (7, '00c27a06-2722-49cd-b233-cde03d3afc7b', '1323c6db-0a94-4b45-bd07-5a7091e777e9');
INSERT INTO finish (pos, race, car_p) VALUES (8, '00c27a06-2722-49cd-b233-cde03d3afc7b', 'f15657ca-919a-420d-8785-f605ce6098ab');
INSERT INTO finish (pos, race, car_p) VALUES (9, '00c27a06-2722-49cd-b233-cde03d3afc7b', '70645b62-8df4-4334-a8d5-79acce411cba');
INSERT INTO finish (pos, race, car_p) VALUES (10, '00c27a06-2722-49cd-b233-cde03d3afc7b', '1461d107-ea45-4dbe-b456-92cba175a824');
INSERT INTO finish (pos, race, car_p) VALUES (11, '00c27a06-2722-49cd-b233-cde03d3afc7b', 'd2335551-2808-47bb-8229-6cd849c3af1a');
INSERT INTO finish (pos, race, car_p) VALUES (12, '00c27a06-2722-49cd-b233-cde03d3afc7b', '9a97bead-1061-4cc7-85a2-866670c59342');
INSERT INTO finish (pos, race, car_p) VALUES (13, '00c27a06-2722-49cd-b233-cde03d3afc7b', 'f334e8a1-28fd-416a-a0ce-dad0423e385b');
INSERT INTO finish (pos, race, car_p) VALUES (14, '00c27a06-2722-49cd-b233-cde03d3afc7b', '9a46a97f-30ad-43a0-aeda-91a280947848');
INSERT INTO finish (pos, race, car_p) VALUES (15, '00c27a06-2722-49cd-b233-cde03d3afc7b', 'b5f023e6-d2c9-4e10-841c-536f5eb176eb');
INSERT INTO finish (pos, race, car_p) VALUES (16, '00c27a06-2722-49cd-b233-cde03d3afc7b', '28d21e2c-ef12-41e4-ab27-3c17b22c60c1');
INSERT INTO finish (pos, race, car_p) VALUES (17, '00c27a06-2722-49cd-b233-cde03d3afc7b', '92116a61-59c8-4b83-ab9e-f4153f3c9337');
INSERT INTO finish (pos, race, car_p) VALUES (18, '00c27a06-2722-49cd-b233-cde03d3afc7b', 'd262ce67-5829-4eb6-befa-48d07e08b6c6');
INSERT INTO finish (pos, race, car_p) VALUES (1, '6b317387-403f-4ad3-9bbc-23cd22fb1310', 'b5f023e6-d2c9-4e10-841c-536f5eb176eb');
INSERT INTO finish (pos, race, car_p) VALUES (2, '6b317387-403f-4ad3-9bbc-23cd22fb1310', 'f15657ca-919a-420d-8785-f605ce6098ab');
INSERT INTO finish (pos, race, car_p) VALUES (3, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '1323c6db-0a94-4b45-bd07-5a7091e777e9');
INSERT INTO finish (pos, race, car_p) VALUES (4, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '716d317e-bde0-4fa3-9738-ec0ef9c3cce1');
INSERT INTO finish (pos, race, car_p) VALUES (5, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '624b63a0-41d9-4c3b-af56-d631afebf840');
INSERT INTO finish (pos, race, car_p) VALUES (6, '6b317387-403f-4ad3-9bbc-23cd22fb1310', 'd2335551-2808-47bb-8229-6cd849c3af1a');
INSERT INTO finish (pos, race, car_p) VALUES (7, '6b317387-403f-4ad3-9bbc-23cd22fb1310', 'c65c3f38-77fd-4099-a10d-35d75d6be870');
INSERT INTO finish (pos, race, car_p) VALUES (8, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '9a97bead-1061-4cc7-85a2-866670c59342');
INSERT INTO finish (pos, race, car_p) VALUES (9, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '1461d107-ea45-4dbe-b456-92cba175a824');
INSERT INTO finish (pos, race, car_p) VALUES (10, '6b317387-403f-4ad3-9bbc-23cd22fb1310', 'd262ce67-5829-4eb6-befa-48d07e08b6c6');
INSERT INTO finish (pos, race, car_p) VALUES (11, '6b317387-403f-4ad3-9bbc-23cd22fb1310', 'c2afd8f8-285f-4903-982a-770a2eacfa04');
INSERT INTO finish (pos, race, car_p) VALUES (12, '6b317387-403f-4ad3-9bbc-23cd22fb1310', 'f334e8a1-28fd-416a-a0ce-dad0423e385b');
INSERT INTO finish (pos, race, car_p) VALUES (13, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '28d21e2c-ef12-41e4-ab27-3c17b22c60c1');
INSERT INTO finish (pos, race, car_p) VALUES (14, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '92116a61-59c8-4b83-ab9e-f4153f3c9337');
INSERT INTO finish (pos, race, car_p) VALUES (15, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '621245bb-31ee-4da9-b224-6e30be870e17');
INSERT INTO finish (pos, race, car_p) VALUES (16, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '9a46a97f-30ad-43a0-aeda-91a280947848');
INSERT INTO finish (pos, race, car_p) VALUES (17, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '1655d35e-8a04-47eb-b694-fcc0f8b8c968');
INSERT INTO finish (pos, race, car_p) VALUES (18, '6b317387-403f-4ad3-9bbc-23cd22fb1310', '70645b62-8df4-4334-a8d5-79acce411cba');
INSERT INTO finish (pos, race, car_p) VALUES (1, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '624b63a0-41d9-4c3b-af56-d631afebf840');
INSERT INTO finish (pos, race, car_p) VALUES (2, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', 'd262ce67-5829-4eb6-befa-48d07e08b6c6');
INSERT INTO finish (pos, race, car_p) VALUES (3, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '1655d35e-8a04-47eb-b694-fcc0f8b8c968');
INSERT INTO finish (pos, race, car_p) VALUES (4, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '9a97bead-1061-4cc7-85a2-866670c59342');
INSERT INTO finish (pos, race, car_p) VALUES (5, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '1323c6db-0a94-4b45-bd07-5a7091e777e9');
INSERT INTO finish (pos, race, car_p) VALUES (6, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '92116a61-59c8-4b83-ab9e-f4153f3c9337');
INSERT INTO finish (pos, race, car_p) VALUES (7, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '716d317e-bde0-4fa3-9738-ec0ef9c3cce1');
INSERT INTO finish (pos, race, car_p) VALUES (8, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '70645b62-8df4-4334-a8d5-79acce411cba');
INSERT INTO finish (pos, race, car_p) VALUES (9, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '621245bb-31ee-4da9-b224-6e30be870e17');
INSERT INTO finish (pos, race, car_p) VALUES (10, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', 'f334e8a1-28fd-416a-a0ce-dad0423e385b');
INSERT INTO finish (pos, race, car_p) VALUES (11, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', 'f15657ca-919a-420d-8785-f605ce6098ab');
INSERT INTO finish (pos, race, car_p) VALUES (12, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '28d21e2c-ef12-41e4-ab27-3c17b22c60c1');
INSERT INTO finish (pos, race, car_p) VALUES (13, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', 'd2335551-2808-47bb-8229-6cd849c3af1a');
INSERT INTO finish (pos, race, car_p) VALUES (14, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', 'c2afd8f8-285f-4903-982a-770a2eacfa04');
INSERT INTO finish (pos, race, car_p) VALUES (15, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '9a46a97f-30ad-43a0-aeda-91a280947848');
INSERT INTO finish (pos, race, car_p) VALUES (16, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', '1461d107-ea45-4dbe-b456-92cba175a824');
INSERT INTO finish (pos, race, car_p) VALUES (17, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', 'b5f023e6-d2c9-4e10-841c-536f5eb176eb');
INSERT INTO finish (pos, race, car_p) VALUES (18, '6cc2306e-75b3-47cd-9c20-fa2ad1a6d964', 'c65c3f38-77fd-4099-a10d-35d75d6be870');


--
-- Data for Name: qualifying; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO qualifying (pos, car_p, race) VALUES (1, 'f9dc7fe7-bc27-4921-b665-96d1aa72ec87', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (2, '3f8291ff-6434-400b-bc28-53a15df6bd18', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (3, '1e1c1666-982b-46f8-9c4a-4350c0994aac', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (4, 'ddff5320-64f6-460b-a76a-19488ea08e70', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (5, '147902c8-1f9c-4ed3-a310-91935ee7b41d', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (6, '8793bec9-4be7-459a-b9da-0dd56a223de6', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (7, 'e03d5b31-1b44-4375-b09c-a303cbbf5d8f', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (8, 'b1d55df9-d7b0-4f95-bd08-ba8009ba1b29', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (9, '60a41976-bb4d-4ce5-ab9c-da8a09106389', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (10, 'd8c0d728-08cb-450a-b5ee-3b98027e80d3', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (11, '44fff04f-d19a-4aab-8ee4-a7cdcfc38fd8', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (12, '011f7d02-b781-443c-bccc-c3d61aab0f24', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (13, 'c76df0ff-11ca-4170-b88a-1291f8442066', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (14, '22984b36-d8fd-48f8-8068-006140dd34da', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (15, '296e1bc6-c1c7-4261-b48a-935986950265', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (16, 'e9da05f6-4335-493e-8b1a-9d050ce33e59', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (17, '00b26904-d63c-45aa-a917-1878c7af767e', '50086d13-1588-4502-b636-91fdaa389383');
INSERT INTO qualifying (pos, car_p, race) VALUES (18, '80923a54-81ab-40e7-b9d5-e7da52d83761', '50086d13-1588-4502-b636-91fdaa389383');


--
-- Data for Name: team_p; Type: TABLE DATA; Schema: public; Owner: -
--

INSERT INTO team_p (car_p, driver) VALUES ('e03d5b31-1b44-4375-b09c-a303cbbf5d8f', '59411c83-b6f3-4724-86dc-c54262166a85');
INSERT INTO team_p (car_p, driver) VALUES ('b1d55df9-d7b0-4f95-bd08-ba8009ba1b29', '5582f399-3a47-4d7e-b325-59fc7888e364');
INSERT INTO team_p (car_p, driver) VALUES ('e03d5b31-1b44-4375-b09c-a303cbbf5d8f', '892a168c-9c19-4356-adfe-ec37308c0ae4');
INSERT INTO team_p (car_p, driver) VALUES ('b1d55df9-d7b0-4f95-bd08-ba8009ba1b29', 'c1f36b6e-47fb-4517-8394-e76021367e50');
INSERT INTO team_p (car_p, driver) VALUES ('b1d55df9-d7b0-4f95-bd08-ba8009ba1b29', 'c86a5bb9-2555-4510-8536-5e3cf3fa7cee');
INSERT INTO team_p (car_p, driver) VALUES ('1e1c1666-982b-46f8-9c4a-4350c0994aac', 'd5f10135-5dc3-4b76-9ca8-9cca944e8eb5');
INSERT INTO team_p (car_p, driver) VALUES ('1e1c1666-982b-46f8-9c4a-4350c0994aac', 'dd32dd80-8842-450f-9d2e-503e18a2505d');
INSERT INTO team_p (car_p, driver) VALUES ('1e1c1666-982b-46f8-9c4a-4350c0994aac', 'ee95cd16-cd67-40a0-90a0-f975bf91d064');
INSERT INTO team_p (car_p, driver) VALUES ('f9dc7fe7-bc27-4921-b665-96d1aa72ec87', 'a8cf93ca-8be8-4bf7-ab24-e568e44e2bc0');
INSERT INTO team_p (car_p, driver) VALUES ('f9dc7fe7-bc27-4921-b665-96d1aa72ec87', '36d5ae0f-c7ed-4c35-9039-4434e2e5f6ff');
INSERT INTO team_p (car_p, driver) VALUES ('f9dc7fe7-bc27-4921-b665-96d1aa72ec87', 'f50c48a9-65e9-4617-bf7a-0e7b0c00b7e3');
INSERT INTO team_p (car_p, driver) VALUES ('e03d5b31-1b44-4375-b09c-a303cbbf5d8f', '13c8ae17-41d7-4321-898d-e82555325a1d');
INSERT INTO team_p (car_p, driver) VALUES ('00b26904-d63c-45aa-a917-1878c7af767e', '8d3e1f53-0290-414c-b788-fd6d13ce1b8d');
INSERT INTO team_p (car_p, driver) VALUES ('00b26904-d63c-45aa-a917-1878c7af767e', '29146774-0472-47a4-8a2e-ead3f93876fc');
INSERT INTO team_p (car_p, driver) VALUES ('00b26904-d63c-45aa-a917-1878c7af767e', 'cce9262f-858b-4318-96ea-e9515441d530');
INSERT INTO team_p (car_p, driver) VALUES ('44fff04f-d19a-4aab-8ee4-a7cdcfc38fd8', '67484669-106c-4f78-b5f5-6605b464b621');
INSERT INTO team_p (car_p, driver) VALUES ('44fff04f-d19a-4aab-8ee4-a7cdcfc38fd8', '84aa0b68-758d-4a91-8ed8-0a42af389125');
INSERT INTO team_p (car_p, driver) VALUES ('44fff04f-d19a-4aab-8ee4-a7cdcfc38fd8', '945da040-e6d2-4544-a36f-a6bf8698b029');
INSERT INTO team_p (car_p, driver) VALUES ('c76df0ff-11ca-4170-b88a-1291f8442066', '2a52185e-77b3-40d0-9065-86a959be6552');
INSERT INTO team_p (car_p, driver) VALUES ('c76df0ff-11ca-4170-b88a-1291f8442066', 'f52f876c-832f-4af0-b6e8-7fc93c047a37');
INSERT INTO team_p (car_p, driver) VALUES ('c76df0ff-11ca-4170-b88a-1291f8442066', '75bf7732-5510-4fb0-8e4d-cb093c04c1bf');
INSERT INTO team_p (car_p, driver) VALUES ('296e1bc6-c1c7-4261-b48a-935986950265', 'e79027fa-44ca-4f3a-b948-6b92931371c2');
INSERT INTO team_p (car_p, driver) VALUES ('296e1bc6-c1c7-4261-b48a-935986950265', 'b2983bcf-c4bb-4678-8215-9ece58afa433');
INSERT INTO team_p (car_p, driver) VALUES ('60a41976-bb4d-4ce5-ab9c-da8a09106389', '2e0cf678-3749-4c9c-8363-a4b585f44a44');
INSERT INTO team_p (car_p, driver) VALUES ('60a41976-bb4d-4ce5-ab9c-da8a09106389', '28dd7ee0-f089-4743-acea-50065e5baf00');
INSERT INTO team_p (car_p, driver) VALUES ('60a41976-bb4d-4ce5-ab9c-da8a09106389', '9e0c8893-5964-45ab-8781-27c0e54f6d1c');
INSERT INTO team_p (car_p, driver) VALUES ('22984b36-d8fd-48f8-8068-006140dd34da', 'd56787ff-9dda-4c5d-a57c-e41567c8cb03');
INSERT INTO team_p (car_p, driver) VALUES ('22984b36-d8fd-48f8-8068-006140dd34da', '69ca8299-d457-4995-b93b-23cad4dd9a2c');
INSERT INTO team_p (car_p, driver) VALUES ('22984b36-d8fd-48f8-8068-006140dd34da', '14120944-7417-4bce-8073-d97c4c348e2d');
INSERT INTO team_p (car_p, driver) VALUES ('d8c0d728-08cb-450a-b5ee-3b98027e80d3', '08dd3fcf-c817-45c5-b50d-4edd52cc0349');
INSERT INTO team_p (car_p, driver) VALUES ('d8c0d728-08cb-450a-b5ee-3b98027e80d3', '2ade3470-e6f3-4ac4-8856-4b74e61b495d');
INSERT INTO team_p (car_p, driver) VALUES ('d8c0d728-08cb-450a-b5ee-3b98027e80d3', '2af51bd9-70e1-414d-b914-6780a8cba253');
INSERT INTO team_p (car_p, driver) VALUES ('011f7d02-b781-443c-bccc-c3d61aab0f24', '595977f4-93cc-453e-ae47-78a51143b040');
INSERT INTO team_p (car_p, driver) VALUES ('011f7d02-b781-443c-bccc-c3d61aab0f24', '2c3cfd0e-4577-4e2a-a51e-e96ab4721bb0');
INSERT INTO team_p (car_p, driver) VALUES ('011f7d02-b781-443c-bccc-c3d61aab0f24', 'd899e2b6-1adf-4cd1-84ea-7dc3412ef6ef');
INSERT INTO team_p (car_p, driver) VALUES ('e9da05f6-4335-493e-8b1a-9d050ce33e59', 'ec1a153f-8e8b-43c4-8cf3-de0351b14995');
INSERT INTO team_p (car_p, driver) VALUES ('e9da05f6-4335-493e-8b1a-9d050ce33e59', 'cda9f220-babf-4002-b7bc-0d00742ba726');
INSERT INTO team_p (car_p, driver) VALUES ('e9da05f6-4335-493e-8b1a-9d050ce33e59', '0554e578-f8ec-4746-95cc-a823aafe97ca');
INSERT INTO team_p (car_p, driver) VALUES ('80923a54-81ab-40e7-b9d5-e7da52d83761', '3cc7bd80-3cfe-4247-8e45-a7e3d1f68b89');
INSERT INTO team_p (car_p, driver) VALUES ('80923a54-81ab-40e7-b9d5-e7da52d83761', '13a33ba6-d793-4712-aced-0e71831f2399');
INSERT INTO team_p (car_p, driver) VALUES ('80923a54-81ab-40e7-b9d5-e7da52d83761', '2c1cb73d-3196-41b2-81a0-74053301bef5');
INSERT INTO team_p (car_p, driver) VALUES ('1655d35e-8a04-47eb-b694-fcc0f8b8c968', 'fdf81545-99a3-4942-bdbf-2d20d0b2c9aa');
INSERT INTO team_p (car_p, driver) VALUES ('1655d35e-8a04-47eb-b694-fcc0f8b8c968', 'e10a6e3e-7886-43ac-93ea-b53ca89257cf');
INSERT INTO team_p (car_p, driver) VALUES ('1655d35e-8a04-47eb-b694-fcc0f8b8c968', 'ad364e55-bab9-4575-becd-21dd849263e9');
INSERT INTO team_p (car_p, driver) VALUES ('624b63a0-41d9-4c3b-af56-d631afebf840', '24477827-d91e-47f5-9706-44659d5c65f2');
INSERT INTO team_p (car_p, driver) VALUES ('624b63a0-41d9-4c3b-af56-d631afebf840', 'e6e630f6-3df8-4be3-aa51-fd2f97d31826');
INSERT INTO team_p (car_p, driver) VALUES ('624b63a0-41d9-4c3b-af56-d631afebf840', '1a19b7fa-ac12-40ce-a014-d78ee255902c');
INSERT INTO team_p (car_p, driver) VALUES ('f334e8a1-28fd-416a-a0ce-dad0423e385b', 'f2496c6f-fff7-4413-ab81-f014766aa871');
INSERT INTO team_p (car_p, driver) VALUES ('f334e8a1-28fd-416a-a0ce-dad0423e385b', '69a376d0-bf40-4550-ae3a-03b9c3d94a26');
INSERT INTO team_p (car_p, driver) VALUES ('f334e8a1-28fd-416a-a0ce-dad0423e385b', 'a1c119e0-cc37-42fb-adec-4fdefb4a0c0a');
INSERT INTO team_p (car_p, driver) VALUES ('716d317e-bde0-4fa3-9738-ec0ef9c3cce1', '694101ca-2a14-4295-a232-aa96999459b3');
INSERT INTO team_p (car_p, driver) VALUES ('716d317e-bde0-4fa3-9738-ec0ef9c3cce1', '81446b3b-50d6-42c6-931f-e78f3daf2585');
INSERT INTO team_p (car_p, driver) VALUES ('716d317e-bde0-4fa3-9738-ec0ef9c3cce1', '2a946a2f-4712-4a76-b94d-26c67543ff75');
INSERT INTO team_p (car_p, driver) VALUES ('e1450469-8b19-4706-a9cb-5ef41fcf500a', '507028e9-bd69-4d0c-a926-3ef0583327aa');
INSERT INTO team_p (car_p, driver) VALUES ('e1450469-8b19-4706-a9cb-5ef41fcf500a', '7908afda-1c81-4c7a-865d-3a6acaf6eada');
INSERT INTO team_p (car_p, driver) VALUES ('e1450469-8b19-4706-a9cb-5ef41fcf500a', '7278b7be-3738-4348-bb35-f60eec9425e4');
INSERT INTO team_p (car_p, driver) VALUES ('bc43d47e-8913-4931-9042-788772e5a9ba', 'f46f108e-bb08-48fc-b504-b01f3d346da8');
INSERT INTO team_p (car_p, driver) VALUES ('bc43d47e-8913-4931-9042-788772e5a9ba', '13ab38a4-a673-4714-a4e8-b2a2798876ab');
INSERT INTO team_p (car_p, driver) VALUES ('bc43d47e-8913-4931-9042-788772e5a9ba', '77ff453d-9312-4696-a1ae-86b79118a440');
INSERT INTO team_p (car_p, driver) VALUES ('c65c3f38-77fd-4099-a10d-35d75d6be870', '29abb75e-f4c7-4169-820e-254e42beb140');
INSERT INTO team_p (car_p, driver) VALUES ('c65c3f38-77fd-4099-a10d-35d75d6be870', 'b309240b-172c-4b66-b3f3-081e8a8fb3b8');
INSERT INTO team_p (car_p, driver) VALUES ('c65c3f38-77fd-4099-a10d-35d75d6be870', 'a2b726ea-f86c-41fe-b72e-683090016665');
INSERT INTO team_p (car_p, driver) VALUES ('621245bb-31ee-4da9-b224-6e30be870e17', '5f68a8f7-d016-4d82-bf7a-8fc2e5665ebb');
INSERT INTO team_p (car_p, driver) VALUES ('621245bb-31ee-4da9-b224-6e30be870e17', '2e3cdac7-8754-42aa-8451-5457d413f1bb');
INSERT INTO team_p (car_p, driver) VALUES ('621245bb-31ee-4da9-b224-6e30be870e17', 'ee9b3b5e-9b17-4169-aa55-cc6c9f4f4e52');
INSERT INTO team_p (car_p, driver) VALUES ('9a46a97f-30ad-43a0-aeda-91a280947848', '17597188-d634-40ab-b901-cbea7095726d');
INSERT INTO team_p (car_p, driver) VALUES ('9a46a97f-30ad-43a0-aeda-91a280947848', '64254f17-7f67-4ea1-8122-b0956c6d5c07');
INSERT INTO team_p (car_p, driver) VALUES ('9a46a97f-30ad-43a0-aeda-91a280947848', 'd11f74e9-15b3-40b1-bd13-969b6efd62f4');
INSERT INTO team_p (car_p, driver) VALUES ('1461d107-ea45-4dbe-b456-92cba175a824', 'de470587-b265-47c1-99ac-fc7fe91eb441');
INSERT INTO team_p (car_p, driver) VALUES ('1461d107-ea45-4dbe-b456-92cba175a824', 'c90a607f-70ff-4113-8682-ad40692db421');
INSERT INTO team_p (car_p, driver) VALUES ('1461d107-ea45-4dbe-b456-92cba175a824', '255809e5-3d8f-4a2f-a5ac-bb773efec9f4');
INSERT INTO team_p (car_p, driver) VALUES ('1323c6db-0a94-4b45-bd07-5a7091e777e9', 'c4b6a6ca-aee8-4367-9fdf-afe3fe6cba4e');
INSERT INTO team_p (car_p, driver) VALUES ('1323c6db-0a94-4b45-bd07-5a7091e777e9', 'a5beded6-7fae-4a55-b6bd-39f72a8d832e');
INSERT INTO team_p (car_p, driver) VALUES ('1323c6db-0a94-4b45-bd07-5a7091e777e9', 'd814cc43-a39a-4e27-be27-61bd506cff97');
INSERT INTO team_p (car_p, driver) VALUES ('92116a61-59c8-4b83-ab9e-f4153f3c9337', 'f8379d1e-a511-4629-96ef-99c418bfdf28');
INSERT INTO team_p (car_p, driver) VALUES ('92116a61-59c8-4b83-ab9e-f4153f3c9337', '0a420565-c6e4-4169-b7a5-2cdbe92e43d2');
INSERT INTO team_p (car_p, driver) VALUES ('296e1bc6-c1c7-4261-b48a-935986950265', 'd1e844b7-5a90-46f1-b7d8-e36f0e5509b0');
INSERT INTO team_p (car_p, driver) VALUES ('ddff5320-64f6-460b-a76a-19488ea08e70', '6840fcc9-fec8-4c8d-9cb8-178b1c96ccef');
INSERT INTO team_p (car_p, driver) VALUES ('ddff5320-64f6-460b-a76a-19488ea08e70', '08f4e629-c708-4b9f-a980-43561d9f1267');
INSERT INTO team_p (car_p, driver) VALUES ('ddff5320-64f6-460b-a76a-19488ea08e70', 'eccc8ea0-17a6-49e1-83ac-b677503c9806');
INSERT INTO team_p (car_p, driver) VALUES ('147902c8-1f9c-4ed3-a310-91935ee7b41d', '9694de26-25b9-4a57-adb4-773023bc2740');
INSERT INTO team_p (car_p, driver) VALUES ('147902c8-1f9c-4ed3-a310-91935ee7b41d', 'c11ab2b7-efdb-4a81-b24a-a4d4e561040e');
INSERT INTO team_p (car_p, driver) VALUES ('147902c8-1f9c-4ed3-a310-91935ee7b41d', '0e82a0b2-8ea0-4569-8b3a-99365e6f44fd');
INSERT INTO team_p (car_p, driver) VALUES ('3f8291ff-6434-400b-bc28-53a15df6bd18', '04a9ad4b-7dea-422c-a9b1-58108285e961');
INSERT INTO team_p (car_p, driver) VALUES ('3f8291ff-6434-400b-bc28-53a15df6bd18', 'f011403a-8b67-47c7-89e0-2798dac2d356');
INSERT INTO team_p (car_p, driver) VALUES ('3f8291ff-6434-400b-bc28-53a15df6bd18', '687eef26-c1c1-4e8b-90b9-b20497f81d09');
INSERT INTO team_p (car_p, driver) VALUES ('8793bec9-4be7-459a-b9da-0dd56a223de6', '8d83f0a6-4b6c-4bc8-8b3e-c4d8a6bffa49');
INSERT INTO team_p (car_p, driver) VALUES ('8793bec9-4be7-459a-b9da-0dd56a223de6', 'cfa420ab-bf0c-48be-bc8d-c8f473ae5a75');
INSERT INTO team_p (car_p, driver) VALUES ('8793bec9-4be7-459a-b9da-0dd56a223de6', 'a4a3225c-d29e-482d-9180-6f17a5943334');
INSERT INTO team_p (car_p, driver) VALUES ('92116a61-59c8-4b83-ab9e-f4153f3c9337', 'd2f1c683-2172-44b7-82a6-3a30d0ed8878');
INSERT INTO team_p (car_p, driver) VALUES ('2ee8a9e4-6224-485d-be98-14b50bd3b608', 'c760613b-c656-41f4-947a-8ec824aeea80');
INSERT INTO team_p (car_p, driver) VALUES ('2ee8a9e4-6224-485d-be98-14b50bd3b608', '1c8402b4-27fa-4596-b7ab-fca1a0c98020');
INSERT INTO team_p (car_p, driver) VALUES ('2ee8a9e4-6224-485d-be98-14b50bd3b608', '9f40e026-bff3-4033-875e-fc886125f6dc');
INSERT INTO team_p (car_p, driver) VALUES ('dd4f8cdc-eb0a-4e93-94cc-72c958729269', '18ff0bf9-19e3-4765-be9b-409e9f8e4aab');
INSERT INTO team_p (car_p, driver) VALUES ('dd4f8cdc-eb0a-4e93-94cc-72c958729269', '9122eaa2-8945-432c-a9c0-bc22d2a0ee76');
INSERT INTO team_p (car_p, driver) VALUES ('dd4f8cdc-eb0a-4e93-94cc-72c958729269', 'ea2f3832-1c51-4b67-a30b-465b37a8da59');
INSERT INTO team_p (car_p, driver) VALUES ('9a97bead-1061-4cc7-85a2-866670c59342', '507028e9-bd69-4d0c-a926-3ef0583327aa');
INSERT INTO team_p (car_p, driver) VALUES ('9a97bead-1061-4cc7-85a2-866670c59342', '7908afda-1c81-4c7a-865d-3a6acaf6eada');
INSERT INTO team_p (car_p, driver) VALUES ('9a97bead-1061-4cc7-85a2-866670c59342', '7278b7be-3738-4348-bb35-f60eec9425e4');
INSERT INTO team_p (car_p, driver) VALUES ('d262ce67-5829-4eb6-befa-48d07e08b6c6', 'f46f108e-bb08-48fc-b504-b01f3d346da8');
INSERT INTO team_p (car_p, driver) VALUES ('d262ce67-5829-4eb6-befa-48d07e08b6c6', '13ab38a4-a673-4714-a4e8-b2a2798876ab');
INSERT INTO team_p (car_p, driver) VALUES ('d262ce67-5829-4eb6-befa-48d07e08b6c6', '77ff453d-9312-4696-a1ae-86b79118a440');
INSERT INTO team_p (car_p, driver) VALUES ('70645b62-8df4-4334-a8d5-79acce411cba', 'a7f771ac-54b2-4d01-9a4b-db05cfa56e80');
INSERT INTO team_p (car_p, driver) VALUES ('70645b62-8df4-4334-a8d5-79acce411cba', '79ab0dc9-4348-449b-80f7-ccb8299695df');
INSERT INTO team_p (car_p, driver) VALUES ('70645b62-8df4-4334-a8d5-79acce411cba', 'd6b8dd33-d6e7-4322-b2a7-ed0b46ab6623');
INSERT INTO team_p (car_p, driver) VALUES ('b5f023e6-d2c9-4e10-841c-536f5eb176eb', '697d287b-bfa9-4fdb-a355-d8972901780e');
INSERT INTO team_p (car_p, driver) VALUES ('b5f023e6-d2c9-4e10-841c-536f5eb176eb', '7bd90695-ae0a-4318-9817-3bf55eb4138e');
INSERT INTO team_p (car_p, driver) VALUES ('b5f023e6-d2c9-4e10-841c-536f5eb176eb', '4ac40f48-e466-4bd7-96cd-c12967998f93');
INSERT INTO team_p (car_p, driver) VALUES ('d2335551-2808-47bb-8229-6cd849c3af1a', 'c7fecf5e-2cc8-4413-add8-4f12d191f64a');
INSERT INTO team_p (car_p, driver) VALUES ('d2335551-2808-47bb-8229-6cd849c3af1a', '3e1175b3-51e1-4c4f-ab35-86f86feec927');
INSERT INTO team_p (car_p, driver) VALUES ('d2335551-2808-47bb-8229-6cd849c3af1a', '617ed046-0e5a-4fb6-b919-4709ee458980');
INSERT INTO team_p (car_p, driver) VALUES ('c2afd8f8-285f-4903-982a-770a2eacfa04', '8be224ad-13e2-4580-9f73-4e87366883c5');
INSERT INTO team_p (car_p, driver) VALUES ('c2afd8f8-285f-4903-982a-770a2eacfa04', '5f372921-dbdd-4f92-b630-7930f6e96cf7');
INSERT INTO team_p (car_p, driver) VALUES ('c2afd8f8-285f-4903-982a-770a2eacfa04', '82e5f157-7eb7-44b3-af7a-020825def80c');
INSERT INTO team_p (car_p, driver) VALUES ('28d21e2c-ef12-41e4-ab27-3c17b22c60c1', 'c760613b-c656-41f4-947a-8ec824aeea80');
INSERT INTO team_p (car_p, driver) VALUES ('28d21e2c-ef12-41e4-ab27-3c17b22c60c1', '1c8402b4-27fa-4596-b7ab-fca1a0c98020');
INSERT INTO team_p (car_p, driver) VALUES ('28d21e2c-ef12-41e4-ab27-3c17b22c60c1', '9f40e026-bff3-4033-875e-fc886125f6dc');
INSERT INTO team_p (car_p, driver) VALUES ('f15657ca-919a-420d-8785-f605ce6098ab', '18ff0bf9-19e3-4765-be9b-409e9f8e4aab');
INSERT INTO team_p (car_p, driver) VALUES ('f15657ca-919a-420d-8785-f605ce6098ab', '9122eaa2-8945-432c-a9c0-bc22d2a0ee76');
INSERT INTO team_p (car_p, driver) VALUES ('f15657ca-919a-420d-8785-f605ce6098ab', 'ea2f3832-1c51-4b67-a30b-465b37a8da59');
INSERT INTO team_p (car_p, driver) VALUES ('77ab202c-f44b-4ba6-b947-84276d9c3769', '6982f216-626b-4c00-8846-5a5cbff8008f');
INSERT INTO team_p (car_p, driver) VALUES ('77ab202c-f44b-4ba6-b947-84276d9c3769', 'f8cae580-7786-47a9-a9b6-44fd0e758264');
INSERT INTO team_p (car_p, driver) VALUES ('77ab202c-f44b-4ba6-b947-84276d9c3769', '54f25a01-a989-4ebe-8a7d-c41315617582');
INSERT INTO team_p (car_p, driver) VALUES ('628c8f62-fb33-42d8-9337-dfc8c0db0d7c', '2327ed6b-c15e-42d9-b41e-3c6a07923e10');
INSERT INTO team_p (car_p, driver) VALUES ('628c8f62-fb33-42d8-9337-dfc8c0db0d7c', '8d020096-bb3a-433e-89ff-9671173d6ea4');
INSERT INTO team_p (car_p, driver) VALUES ('628c8f62-fb33-42d8-9337-dfc8c0db0d7c', '8d5ff1cb-3059-4420-9b76-2e2f189349ac');


--
-- PostgreSQL database dump complete
--

