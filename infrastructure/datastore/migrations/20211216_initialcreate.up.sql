BEGIN;

CREATE TABLE beer (
    id uuid PRIMARY KEY NOT NULL,
    name varchar(100) NOT NULL,
    brewery varchar(100) NOT NULL,
    country varchar(100) NOT NULL,
    price  float,
    currency varchar(100) NOT NULL 
);



/*TODO: ADD A COUPLE OF BEERS c: */
/*INSERT INTO public.beer(
	id, name, brewery, country, price, currency)
	VALUES (?, ?, ?, ?, ?, ?);/
*/
COMMIT;
