
CREATE TABLE IF NOT EXISTS member_post_list (
                                memberid int4 NOT NULL,
                                postid uuid NOT NULL UNIQUE ,
                                CONSTRAINT uk_imgg0icxjiwr4sa7fy1g3eigm UNIQUE (postid)
);

ALTER TABLE public.member_post_list ADD CONSTRAINT fkajghwc1uw1gxjbligrnm3bslk FOREIGN KEY (memberid) REFERENCES member(id);
ALTER TABLE public.member_post_list ADD CONSTRAINT fku1340c6j8ncerdv20dagksp7 FOREIGN KEY (postid) REFERENCES post(id);
