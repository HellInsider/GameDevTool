-- This script was generated by a beta version of the ERD tool in pgAdmin 4.
-- Please log an issue at https://redmine.postgresql.org/projects/pgadmin4/issues/new if you find any bugs, including reproduction steps.
BEGIN;

CREATE SCHEMA "DB_schema";
CREATE TABLE IF NOT EXISTS "DB_schema"."Games"
(
    appid integer NOT NULL,
    name text COLLATE pg_catalog."default" NOT NULL,
    required_age integer,
    is_free boolean,
    dlc integer[],
    about_the_game text COLLATE pg_catalog."default",
    detailed_description text COLLATE pg_catalog."default",
    short_description text COLLATE pg_catalog."default",
    developers text[] COLLATE pg_catalog."default",
    publishers text[] COLLATE pg_catalog."default",
    packages integer[],
    recommendations integer,
    type text DEFAULT 'game',
    requirements_id integer,
    critic_score_id integer,
    release_date_id integer,
    platforms_id integer,
    CONSTRAINT "Games_pkey" PRIMARY KEY (appid)
);

CREATE TABLE IF NOT EXISTS "DB_schema"."Genre"
(
    genre_id integer NOT NULL,
    description text NOT NULL,
    PRIMARY KEY (genre_id)
);

CREATE TABLE IF NOT EXISTS "DB_schema"."Requirements"
(
    appid integer NOT NULL,
    minimum text,
    requirement_id integer NOT NULL,
    recommended text,
    PRIMARY KEY (requirement_id),
    UNIQUE (appid)
);

CREATE TABLE IF NOT EXISTS "DB_schema"."CriticScore"
(
    score_id integer NOT NULL,
    critic_score integer,
    url text,
    user_score integer,
    appid integer NOT NULL,
    PRIMARY KEY (score_id),
    UNIQUE (appid)
);

CREATE TABLE IF NOT EXISTS "DB_schema"."Category"
(
    category_id integer NOT NULL,
    description text NOT NULL,
    PRIMARY KEY (category_id)
);

CREATE TABLE IF NOT EXISTS "DB_schema"."Release_date"
(
    appid integer NOT NULL,
    comming_soon boolean DEFAULT true,
    date text,
    rdate_id integer NOT NULL,
    PRIMARY KEY (rdate_id),
    UNIQUE (appid)
);

CREATE TABLE IF NOT EXISTS "DB_schema"."Platforms"
(
    appid integer NOT NULL,
    platform_win boolean NOT NULL DEFAULT false,
    platform_mac boolean NOT NULL DEFAULT false,
    platform_linux boolean NOT NULL DEFAULT false,
    platform_id integer NOT NULL,
    PRIMARY KEY (platform_id),
    UNIQUE (appid)
);

CREATE TABLE IF NOT EXISTS "DB_schema"."User"
(
    user_id integer NOT NULL,
    name text NOT NULL,
    games_count integer NOT NULL,
    PRIMARY KEY (user_id)
);

CREATE TABLE IF NOT EXISTS "DB_schema"."Played_game"
(
    appid integer NOT NULL,
    playtime_forever integer,
    playtime_windows_forever integer,
    playtime_mac_forever integer,
    playtime_linux_forever integer,
    rtime_last_played integer,
    playtime_2weeks integer,
    has_community_visible_stats boolean,
    PRIMARY KEY (appid)
);

CREATE TABLE IF NOT EXISTS "DB_schema"."Games_Genre"
(
    "Games_appid" integer NOT NULL,
    "Genre_genre_id" integer NOT NULL,
    PRIMARY KEY ("Games_appid", "Genre_genre_id")
);

CREATE TABLE IF NOT EXISTS "DB_schema"."Games_Category"
(
    "Games_appid" integer NOT NULL,
    "Category_category_id" integer NOT NULL,
    PRIMARY KEY ("Games_appid", "Category_category_id")
);

CREATE TABLE IF NOT EXISTS "DB_schema"."User_Played_game"
(
    "User_user_id" integer NOT NULL,
    "Played_game_appid" integer NOT NULL,
    PRIMARY KEY ("User_user_id", "Played_game_appid")
);

CREATE TABLE IF NOT EXISTS "DB_schema"."Prices"
(
    price_id bigserial NOT NULL,
    currency text NOT NULL DEFAULT 'RUB',
    initial integer,
    final integer,
    discount_percent integer,
    country text DEFAULT 'Russia',
    appid integer NOT NULL,
    PRIMARY KEY (price_id)
);

CREATE TABLE IF NOT EXISTS "DB_schema"."Games_Prices"
(
    "Games_appid" integer NOT NULL,
    "Prices_price_id" bigserial NOT NULL,
    PRIMARY KEY ("Games_appid", "Prices_price_id")
);

CREATE TABLE IF NOT EXISTS "DB_schema"."Reviews"
(
    appid integer NOT NULL,
    total_reviews integer DEFAULT 0,
    total_positive integer DEFAULT 0,
    total_negative integer DEFAULT 0,
    review_score integer DEFAULT 0,
    review_score_desc text,
    review_id integer NOT NULL,
    PRIMARY KEY (review_id),
    UNIQUE (appid)
);

ALTER TABLE IF EXISTS "DB_schema"."Games"
    ADD FOREIGN KEY (critic_score_id)
    REFERENCES "DB_schema"."CriticScore" (score_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    DEFERRABLE INITIALLY DEFERRED
    NOT VALID;


ALTER TABLE IF EXISTS "DB_schema"."Games"
    ADD FOREIGN KEY (requirements_id)
    REFERENCES "DB_schema"."Requirements" (requirement_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    DEFERRABLE INITIALLY DEFERRED
    NOT VALID;


ALTER TABLE IF EXISTS "DB_schema"."Games"
    ADD FOREIGN KEY (release_date_id)
    REFERENCES "DB_schema"."Release_date" (rdate_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    DEFERRABLE INITIALLY DEFERRED
    NOT VALID;


ALTER TABLE IF EXISTS "DB_schema"."Games"
    ADD FOREIGN KEY (platforms_id)
    REFERENCES "DB_schema"."Platforms" (platform_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    DEFERRABLE INITIALLY DEFERRED
    NOT VALID;


ALTER TABLE IF EXISTS "DB_schema"."Games"
    ADD COLUMN reviews_id integer,
    ADD FOREIGN KEY (reviews_id)
    REFERENCES "DB_schema"."Reviews" (review_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS "DB_schema"."Played_game"
    ADD FOREIGN KEY (appid)
    REFERENCES "DB_schema"."Games" (appid) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS "DB_schema"."Games_Genre"
    ADD FOREIGN KEY ("Games_appid")
    REFERENCES "DB_schema"."Games" (appid) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS "DB_schema"."Games_Genre"
    ADD FOREIGN KEY ("Genre_genre_id")
    REFERENCES "DB_schema"."Genre" (genre_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS "DB_schema"."Games_Category"
    ADD FOREIGN KEY ("Games_appid")
    REFERENCES "DB_schema"."Games" (appid) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS "DB_schema"."Games_Category"
    ADD FOREIGN KEY ("Category_category_id")
    REFERENCES "DB_schema"."Category" (category_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS "DB_schema"."User_Played_game"
    ADD FOREIGN KEY ("User_user_id")
    REFERENCES "DB_schema"."User" (user_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS "DB_schema"."User_Played_game"
    ADD FOREIGN KEY ("Played_game_appid")
    REFERENCES "DB_schema"."Played_game" (appid) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS "DB_schema"."Games_Prices"
    ADD FOREIGN KEY ("Games_appid")
    REFERENCES "DB_schema"."Games" (appid) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;


ALTER TABLE IF EXISTS "DB_schema"."Games_Prices"
    ADD FOREIGN KEY ("Prices_price_id")
    REFERENCES "DB_schema"."Prices" (price_id) MATCH SIMPLE
    ON UPDATE NO ACTION
    ON DELETE NO ACTION
    NOT VALID;

END;


CREATE OR REPLACE FUNCTION public.addPrice(_appid integer, _currency text, _initial integer, _final integer, _discount_percent integer, _country text)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
DECLARE p_id integer;
BEGIN
    SELECT "Prices_price_id" INTO p_id
    FROM "DB_schema"."Games_Prices" AS gp
	JOIN "DB_schema"."Prices" AS p ON gp."Games_appid"=p."appid"
    WHERE p.country = _country;
IF EXISTS (SELECT "Prices_price_id"
    FROM "DB_schema"."Games_Prices" AS gp
	JOIN "DB_schema"."Prices" AS p ON gp."Games_appid"=p."appid"
    WHERE p.country = _country) 
THEN
    UPDATE "DB_schema"."Prices" 
    SET currency = _currency, initial = _initial, final = _final, discount_percent=_discount_percent
    WHERE (price_id = p_id);
ELSE    
    INSERT INTO "DB_schema"."Prices" 
    ( currency, initial, final, discount_percent, country, appid) 
    VALUES ( _currency, _initial, _final, _discount_percent, _country, _appid)
    RETURNING price_id INTO p_id;
    
    INSERT INTO "DB_schema"."Games_Prices"
    VALUES (_appid, p_id);	
END IF;
END;
$function$;


CREATE OR REPLACE FUNCTION public.addCategory(_category_id integer, _description text, _appid integer)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
BEGIN
INSERT INTO "DB_schema"."Category" (category_id, description)
VALUES (_category_id, _description) ON CONFLICT DO NOTHING;

INSERT INTO "DB_schema"."Games_Category" ("Games_appid", "Category_category_id") 
VALUES(_appid, _category_id) ON CONFLICT DO NOTHING;
END;
$function$
;

CREATE OR REPLACE FUNCTION public.addGenre(genre_id integer, description text, appid integer)
 RETURNS void
 LANGUAGE plpgsql
AS $function$
BEGIN
INSERT INTO "DB_schema"."Genre"("genre_id", "description")
VALUES(genre_id, description) ON CONFLICT DO NOTHING;

INSERT INTO "DB_schema"."Games_Genre" ("Games_appid", "Genre_genre_id") 
VALUES(appid, genre_id) ON CONFLICT DO NOTHING;
END;
$function$;


END;