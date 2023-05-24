package SQLFunctions

var AddBaseGameInfo = `INSERT INTO "DB_schema"."Games" (appid, name) 
VALUES ($1, $2) ON CONFLICT DO NOTHING;`

var UpdateGameDetails = `UPDATE "DB_schema"."Games" 
SET required_age = $2, is_free = $3, dlc = $4, about_the_game = $5, detailed_description = $6, 
short_description = $7, developers = $8, publishers = $9, packages = $10, recommendations = $11, type = $12,
requirements_id = $1, critic_score_id = $1, release_date_id = $1, platforms_id = $1
WHERE appid = $1`

//VALUES ($2, $3, $4, $5, $6 , $7, $8, $9, $10 , $11, $12, $13)

var AddPrice = `
IF EXISTS (SELECT "Prices_price_id" AS temp_price_id
    FROM "DB_schema"."Games_Prices" AS gp
	JOIN "DB_schema"."Prices" AS p ON gp."Games_appid"=p."appid"
    WHERE p.country = _country ) 
THEN
    UPDATE "DB_schema"."Prices" 
    SET currency = _currence, initial = _initial, final = _final, discount_percent=_discount_percent
    WHERE (price_id = temp_price_id);
ELSE    
    INSERT INTO "DB_schema"."Prices" 
    ( currency, initial, final, discount_percent, country) 
    VALUES ( _currency, _initial, _final, _discount_percent, _country);
    newid = currval('price_id');
    INSERT INTO "DB_schema"."Games_Prices"
    VALUES (_appid, newid);	
END IF;
`

var AddReleaseDate = `INSERT INTO  "DB_schema"."Release_date" 
(appid, comming_soon, date, rdate_id)
VALUES ($1, $2, $3, $1) ON CONFLICT (rdate_id) DO UPDATE
SET  appid =$1, comming_soon = $2, date = $3`

var AddReviews = `INSERT INTO  "DB_schema"."Reviews" 
(appid, total_reviews, total_positive, total_negative, review_score, review_score_desc, review_id)
VALUES ($1, $2, $3, $4, $5, $6, $1) ON CONFLICT (review_id) DO UPDATE
SET  appid =$1, total_reviews = $2, total_positive = $3,total_negative = $4, review_score=$5, review_score_desc=$6`

var AddRequirements = `INSERT INTO  "DB_schema"."Requirements"
(appid, minimum, recommended, requirement_id)
VALUES ($1, $2, $3, $1) ON CONFLICT (requirement_id) DO UPDATE
SET appid =$1, minimum = $2, recommended = $3`

var AddPlatforms = `INSERT INTO  "DB_schema"."Platforms" 
(appid, platform_win, platform_mac, platform_linux, platform_id)
VALUES ($1, $2, $3, $4, $1) ON CONFLICT (platform_id) DO UPDATE
SET  appid = $1, platform_win = $2, platform_mac = $3, platform_linux =$4`

var AddCriticScore = `INSERT INTO  "DB_schema"."CriticScore" 
(appid, critic_score, user_score, url, score_id)
VALUES ($1, $2, $3, $4, $1) ON CONFLICT (score_id) DO UPDATE
SET appid =$1, critic_score = $2, user_score = $3, url = $4`

var AddGenre = `
INSERT INTO "DB_schema"."Genre"(genre_id, description)
VALUES($1, $2) ON CONFLICT DO NOTHING;

INSERT INTO "DB_schema"."Games_Genre" (Games_appid, Genre_genre_id) 
VALUES($3, $1) ON CONFLICT DO NOTHING;`

var AddCategory = `
INSERT INTO "DB_schema"."Category" (category_id, description)
VALUES ($1, $2) ON CONFLICT DO NOTHING;

INSERT INTO "DB_schema"."Games_Category" (Games_appid, Category_category_id) 
VALUES($3, $1) ON CONFLICT DO NOTHING;`

/*
var AddGenre = `
IF NOT EXISTS (SELECT * FROM "DB_schema"."Genre"
                   WHERE genre_id = &1)
BEGIN
INSERT INTO "DB_schema"."Genre" (genre_id, description)
VALUES($1, $2)
END
`

var AddCategory = `
IF NOT EXISTS (SELECT * FROM "DB_schema"."Category"
                   WHERE category_id = &1)
BEGIN
INSERT INTO "DB_schema"."Category" (category_id, description)
VALUES($1, $2)
END
`
*/
