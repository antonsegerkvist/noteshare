insert into t_user (
	c_email,
    c_username,
    c_password_hash,
    c_password_salt,
    c_activated
)
values (
	'root@noteshare.com',
    'root',
    'A513F966BF8C289E511B1D29AEDD6D652D550E73F31A30C52EB20267EB8DF1FF',
    'CEF358063EBA83AFD22AEC05A3A97B106BFAD27E8E6158F3266196E08CBA4DDD',
    NOW()
);