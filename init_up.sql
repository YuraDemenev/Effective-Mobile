--
-- PostgreSQL database dump
--

-- Dumped from database version 16.4
-- Dumped by pg_dump version 16.4

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: groups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.groups (
    id bigint NOT NULL,
    name character varying(100) NOT NULL
);


ALTER TABLE public.groups OWNER TO postgres;

--
-- Name: groups_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.groups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.groups_id_seq OWNER TO postgres;

--
-- Name: groups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.groups_id_seq OWNED BY public.groups.id;


--
-- Name: songs; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.songs (
    id bigint NOT NULL,
    group_id integer NOT NULL,
    name character varying(100),
    link text,
    release_date character varying(20)
);


ALTER TABLE public.songs OWNER TO postgres;

--
-- Name: songs_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.songs_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.songs_id_seq OWNER TO postgres;

--
-- Name: songs_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.songs_id_seq OWNED BY public.songs.id;


--
-- Name: verses; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.verses (
    id integer NOT NULL,
    song_id integer,
    verse_number integer NOT NULL,
    text text NOT NULL
);


ALTER TABLE public.verses OWNER TO postgres;

--
-- Name: verses_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.verses_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.verses_id_seq OWNER TO postgres;

--
-- Name: verses_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.verses_id_seq OWNED BY public.verses.id;


--
-- Name: groups id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.groups ALTER COLUMN id SET DEFAULT nextval('public.groups_id_seq'::regclass);


--
-- Name: songs id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.songs ALTER COLUMN id SET DEFAULT nextval('public.songs_id_seq'::regclass);


--
-- Name: verses id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.verses ALTER COLUMN id SET DEFAULT nextval('public.verses_id_seq'::regclass);


--
-- Data for Name: groups; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.groups (id, name) FROM stdin;
1	Red Hot Chili Peppers
2	Gorillaz
3	Nirvana
\.


--
-- Data for Name: songs; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.songs (id, group_id, name, link, release_date) FROM stdin;
1	1	Californication	https://www.youtube.com/watch?v=OtXiwSCq99Q	08.06.1999
2	1	Dani California	https://www.youtube.com/watch?v=Sb5aq5HcS1A&list=RDYlUKcNNmywk&index=7	03.04.2006
3	1	Dark Necessities	https://www.youtube.com/watch?v=Q0oIoR9mLwc	05.03.2016
4	2	Clint Eastwood	https://www.youtube.com/watch?v=1V_xRb0x9aw	05.03.2001
5	3	Smells Like Teen Spirit	https://www.youtube.com/watch?v=hTWKbfoikeg	10.09.1991
\.


--
-- Data for Name: verses; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.verses (id, song_id, verse_number, text) FROM stdin;
1	1	1	Psychic spies from China try to steal your mind's elation And little girls from Sweden dream of silver-screen quotation And if you want these kind of dreams, it's Californication 
2	1	2	 It's the edge of the world and all of Western civilization The sun may rise in the East, at least it settled in a final location It's understood that Hollywood sells Californication 
3	1	3	Pay your surgeon very well to break the spell of aging Celebrity skin: is this your chin or is that war you're waging? First born unicorn Hardcore soft porn 
4	1	4	Dream of Californication Dream of Californication Dream of Californication Dream of Californication 
5	1	5	Marry me, girl, be my fairy to the world, be my very own constellation A teenage bride with a baby inside getting high on information And buy me a star on the boulevard, it's Californication See rock shows near Amsterdam Get tickets as low as $22 You might also like The Alchemy Taylor Swift 6:16 in LA 
6	2	1	[Verse 1: Anthony Kiedis] Gettin' born in the state of Mississippi Poppa was a copper and her momma was a hippie In Alabama, she would swing a hammer Price you gotta pay when you break the panorama 
7	2	2	[Pre-Chorus: Anthony Kiedis] She never knew that there was anything more than poor What in the world does your company take me for? 
8	2	3	[Verse 2: Anthony Kiedis] Black bandana, sweet Louisiana Robbin' a bank in the state of Indiana She's a runner, rebel and a stunner On her merry way sayin', Baby, what you gonna—? 
9	2	4	[Pre-Chorus: Anthony Kiedis] Lookin' down the barrel of a hot metal .45 Just another way to survive 
10	2	5	[Chorus: Anthony Kiedis] California, rest in peace Simultaneous release California, show your teeth She's my priestess, I'm your priest, yeah, yeah See rock shows near Amsterdam Get tickets as low as $22 You might also like So Long, London Taylor Swift The Alchemy Taylor Swift 7 Minute Drill J. Cole 
11	2	6	[Verse 3: Anthony Kiedis] She's a lover, baby and a fighter Shoulda seen her comin' when it got a little brighter With a name like Dani California Day was gonna come when I was gonna mourn ya 
12	2	7	[Pre-Chorus: Anthony Kiedis] A little loaded, she was stealin' another breath I love my baby to death 
13	2	8	[Chorus: Anthony Kiedis] California, rest in peace Simultaneous release California, show your teeth She's my priestess, I'm your priest, yeah, yeah 
14	2	9	[Bridge: John Frusciante] Who knew the other side of you? Who knew what others died to prove? Too true to say goodbye to you Too true to say, say, say 
15	2	10	[Verse 4: Anthony Kiedis] Push the fader, gifted animator One for the now and eleven for the later Never made it up to Minnesota North Dakota man was a-gunnin' for the quota [Pre-Chorus: Anthony Kiedis] Down in the Badlands, she was savin' the best for last It only hurts when I laugh Gone too fast 
16	2	11	[Chorus: Anthony Kiedis] California, rest in peace Simultaneous release California, show your teeth She's my priestess, I'm your priest, yeah, yeah California, rest in peace (Do svidaniya) Simultaneous release (California) California, show your teeth (Do svidaniya) She's my priestess, I'm your priest, yeah, yeah
17	3	1	[Verse 1] Comin' on to the light of day, we got Many moons that are deep at play so I Keep an eye on the shadow smile to see what it has to say You and I both know, everything must go away Ah, what do you say? Spinnin' knot that is on my heart is like a Bit of light in a touch of dark, you got Sneak attack from the zodiac but I see your fire spark Eat the breeze and go, blow by blow and go away Oh, what do you say? Yeah 
18	3	2	[Chorus] You don't know my mind, you don't know my kind Dark necessities are part of my design and Tell the world that I'm falling from the sky Dark necessities are part of my design 
19	3	3	[Verse 2] Stumble down to the parking lot, you got No time for the afterthought, they're like Ice Cream for an Astronaut, well, that's me looking for we Turn the corner and find the world at your command Playin' the hand, yeah 
20	3	4	[Chorus] You don't know my mind, you don't know my kind Dark necessities are part of my design Tell the world that I'm falling from the sky Dark necessities are part of my design 
21	3	5	[Bridge] Do you want this love of mine? Darkness helps us all to shine Do you want it, do you want it now? Do you want it all the time? But darkness helps us all to shine Do you want it, do you want it now? 
22	3	6	[Verse 3] Ah, pick you up like a paperback with the Track record of a maniac so I Move it in and we unpack, it's the same as yesterday Any way we roll, everything must go away Oh, what do you say? Yeah 
23	3	7	[Chorus] You don't know my mind, you don't know my kind  Dark necessities are part of my design Tell the world that I'm falling from the sky and Dark necessities are part of my design 
24	3	8	[Outro] Ah-ah-ah Ah-ah Ah-ah-ah Ah-ah Ah-ah-ah Ah-ah Ah-ah-ah Ah-ah
25	4	1	[Intro: 2-D] Hoo-hoo-hoo-hoo-hoo 
26	4	2	[Chorus: 2-D] I ain't happy, I'm feeling glad I got sunshine in a bag I'm useless, but not for long The future is coming on 
27	4	3	I ain't happy, I'm feeling glad I got sunshine in a bag I'm useless, but not for long The future is coming on It's coming on, it's coming on It's coming on, it's coming on 
28	4	4	[Verse 1: Del the Funky Homosapien] (Hoo, yeah, haha) Finally, someone let me out of my cage Now, time for me is nothing, 'cause I'm countin' no age Nah, I couldn't be there, now you shouldn't be scared I'm good at repairs, and I'm under each snare Intangible, bet you didn't think, so I command you to Panoramic view, look, I'll make it all manageable Pick and choose, sit and lose, all you different crews Chicks and dudes, who you think is really kickin' tunes? Picture you getting down in a picture tube Like you lit the fuse, you think it's fictional? Mystical? Maybe, spiritual hero 
29	4	5	[Chorus: 2-D & (Del the Funky Homosapien)] I ain't happy, I'm feeling glad I got sunshine in a bag I'm useless, but not for long The future is coming on I ain't happy, I'm feeling glad I got sunshine in a bag I'm useless, but not for long The future is coming on (That's right) It's coming on, it's coming on It's coming on, it's coming on 
30	4	6	[Verse 2: Del the Funky Homosapien] The essence, the basics, without it, you make it Allow me to make this childlike in nature Rhythm, you have it or you don't, that's a fallacy I'm in them, every sproutin' tree, every child of peace Every cloud and sea, you see with your eyes I see destruction and demise, corruption in disguise (That's right) From this fuckin' enterprise, now I'm sucked into your lies Through Russel, not his muscles But percussion he provides for me as a guide, y'all can see me now 'Cause you don't see with your eye, you perceive with your mind That's the inner (Fuck 'em) so I'ma stick around with Russ and be a mentor Bust a few rhymes so motherfuckers remember Where the thought is, I brought all this So you can survive when law is lawless (Right here) Feelings, sensations that you thought was dead No squealing and remember that it's all in your head 
31	4	7	[Chorus: 2-D] I ain't happy, I'm feeling glad I got sunshine in a bag I'm useless, but not for long The future is coming on I ain't happy, I'm feeling glad I got sunshine in a bag I'm useless, but not for long My future is coming on It's coming on, it's coming on It's coming on, it's coming on 
32	4	8	[Outro: 2-D] My future is coming on It's coming on, it's coming on It's coming on, it's coming on My future is coming on It's coming on, it's coming on It's coming on, it's coming on My future is coming on It's coming on, it's coming on My future is coming on It's coming on, it's coming on My future is coming on It's coming on, it's coming on My future
33	5	1	[Verse 1] Load up on guns, bring your friends It's fun to lose and to pretend She's over-bored and self-assured Oh no, I know a dirty word 
34	5	2	[Pre-Chorus]Hello, hello, hello, how low Hello, hello, hello, how low Hello, hello, hello, how low Hello, hello, hello 
35	5	3	[Chorus] With the lights out, it's less dangerous Here we are now, entertain us I feel stupid and contagious Here we are now, entertain us A mulatto, an albino A mosquito, my libido, yeah 
36	5	4	[Post-Chorus] Hey, yay 
37	5	5	[Verse 2] I'm worse at what I do best  And for this gift, I feel blessed Our little group has always been And always will until the end 
38	5	6	[Pre-Chorus]Hello, hello, hello, how low Hello, hello, hello, how low Hello, hello, hello, how low Hello, hello, hello 
39	5	7	[Chorus] With the lights out, it's less dangerous Here we are now, entertain us I feel stupid and contagious Here we are now, entertain us A mulatto, an albino A mosquito, my libido, yeah 
40	5	8	[Post-Chorus] Hey, yay 
41	5	9	[Verse 3] And I forget just why I taste Oh yeah, I guess it makes me smile I found it hard, it's hard to find Oh well, whatever, never mind 
42	5	10	[Pre-Chorus] Hello, hello, hello, how low Hello, hello, hello, how low Hello, hello, hello, how low Hello, hello, hello 
43	5	11	[Chorus] With the lights out, it's less dangerous Here we are now, entertain us I feel stupid and contagious Here we are now, entertain us A mulatto, an albino A mosquito, my libido 
44	5	12	[Outro] A denial, a denial A denial, a denial A denial, a denial A denial, a denial A denial
\.


--
-- Name: groups_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.groups_id_seq', 5, true);


--
-- Name: songs_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.songs_id_seq', 7, true);


--
-- Name: verses_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.verses_id_seq', 46, true);


--
-- Name: groups groups_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_name_key UNIQUE (name);


--
-- Name: groups groups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.groups
    ADD CONSTRAINT groups_pkey PRIMARY KEY (id);


--
-- Name: songs songs_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.songs
    ADD CONSTRAINT songs_pkey PRIMARY KEY (id);


--
-- Name: verses verses_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.verses
    ADD CONSTRAINT verses_pkey PRIMARY KEY (id);


--
-- Name: idx_groups_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_groups_name ON public.groups USING btree (name);


--
-- Name: idx_songs_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_songs_id ON public.songs USING btree (id);


--
-- Name: idx_verses_number; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_verses_number ON public.verses USING btree (song_id, verse_number);


--
-- Name: songs songs_group_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.songs
    ADD CONSTRAINT songs_group_id_fkey FOREIGN KEY (group_id) REFERENCES public.groups(id);


--
-- Name: verses verses_song_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.verses
    ADD CONSTRAINT verses_song_id_fkey FOREIGN KEY (song_id) REFERENCES public.songs(id) ON DELETE CASCADE;


--
-- PostgreSQL database dump complete
--

