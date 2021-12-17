BEGIN;

CREATE TABLE beer (
    id uuid PRIMARY KEY NOT NULL,
    name varchar(100) NOT NULL,
    brewery varchar(100) NOT NULL,
    country varchar(100) NOT NULL,
    price  float,
    currency varchar(100) NOT NULL 
);

    INSERT INTO public.beer(
	id, name, brewery, country, price, currency)
	VALUES ('2bcaaa15-a7c0-4d57-bafa-b5d6c2cf410d', 'Beer Cool Name', 'ColumbiaBeerCO', 'Colombia', 1.2, 'USD');

    INSERT INTO public.beer(
	id, name, brewery, country, price, currency)
	VALUES ('5f21ed03-513a-4731-9323-59d46c1d739b', 'Vusenwaiser', 'GermanyCO', 'Germany', 3.2, 'USD');

    INSERT INTO public.beer(
	id, name, brewery, country, price, currency)
	VALUES ('56dfa6bc-5b13-4401-a4d0-15e1fff3a784', 'Nice Germany Name', 'Germany Company', 'France', 5.2, 'USD');

COMMIT;
