create table if not exists geoip_asn_blocks_ipv4
(
    prefix                         String,
    autonomous_system_number       UInt32,
    autonomous_system_organization String
) ENGINE = Dictionary('geoip_asn_blocks_ipv4');

create table if not exists geoip_asn_blocks_ipv6
(
    prefix                         String,
    autonomous_system_number       UInt32,
    autonomous_system_organization String
) ENGINE = Dictionary('geoip_asn_blocks_ipv6');

create table if not exists geoip_city_locations_de
(
    geoname_id             UInt64,
    locale_code            String,
    continent_code         String,
    continent_name         String,
    country_iso_code       String,
    country_name           String,
    subdivision_1_iso_code String,
    subdivision_1_name     String,
    subdivision_2_iso_code String,
    subdivision_2_name     String,
    city_name              String,
    metro_code             UInt32,
    time_zone              String,
    is_in_european_union   UInt8
) ENGINE = Dictionary('geoip_city_locations_de');

create table if not exists geoip_city_locations_en
(
    geoname_id             UInt64,
    locale_code            String,
    continent_code         String,
    continent_name         String,
    country_iso_code       String,
    country_name           String,
    subdivision_1_iso_code String,
    subdivision_1_name     String,
    subdivision_2_iso_code String,
    subdivision_2_name     String,
    city_name              String,
    metro_code             UInt32,
    time_zone              String,
    is_in_european_union   UInt8
) ENGINE = Dictionary('geoip_city_locations_en');

create table if not exists geoip_city_locations_es
(
    geoname_id             UInt64,
    locale_code            String,
    continent_code         String,
    continent_name         String,
    country_iso_code       String,
    country_name           String,
    subdivision_1_iso_code String,
    subdivision_1_name     String,
    subdivision_2_iso_code String,
    subdivision_2_name     String,
    city_name              String,
    metro_code             UInt32,
    time_zone              String,
    is_in_european_union   UInt8
) ENGINE = Dictionary('geoip_city_locations_es');

create table if not exists geoip_city_locations_fr
(
    geoname_id             UInt64,
    locale_code            String,
    continent_code         String,
    continent_name         String,
    country_iso_code       String,
    country_name           String,
    subdivision_1_iso_code String,
    subdivision_1_name     String,
    subdivision_2_iso_code String,
    subdivision_2_name     String,
    city_name              String,
    metro_code             UInt32,
    time_zone              String,
    is_in_european_union   UInt8
) ENGINE = Dictionary('geoip_city_locations_fr');

create table if not exists geoip_city_locations_ja
(
    geoname_id             UInt64,
    locale_code            String,
    continent_code         String,
    continent_name         String,
    country_iso_code       String,
    country_name           String,
    subdivision_1_iso_code String,
    subdivision_1_name     String,
    subdivision_2_iso_code String,
    subdivision_2_name     String,
    city_name              String,
    metro_code             UInt32,
    time_zone              String,
    is_in_european_union   UInt8
) ENGINE = Dictionary('geoip_city_locations_ja');

create table if not exists geoip_city_locations_pt
(
    geoname_id             UInt64,
    locale_code            String,
    continent_code         String,
    continent_name         String,
    country_iso_code       String,
    country_name           String,
    subdivision_1_iso_code String,
    subdivision_1_name     String,
    subdivision_2_iso_code String,
    subdivision_2_name     String,
    city_name              String,
    metro_code             UInt32,
    time_zone              String,
    is_in_european_union   UInt8
) ENGINE = Dictionary('geoip_city_locations_pt');

create table if not exists geoip_city_locations_ru
(
    geoname_id             UInt64,
    locale_code            String,
    continent_code         String,
    continent_name         String,
    country_iso_code       String,
    country_name           String,
    subdivision_1_iso_code String,
    subdivision_1_name     String,
    subdivision_2_iso_code String,
    subdivision_2_name     String,
    city_name              String,
    metro_code             UInt32,
    time_zone              String,
    is_in_european_union   UInt8
) ENGINE = Dictionary('geoip_city_locations_ru');

create table if not exists geoip_city_locations_zh
(
    geoname_id             UInt64,
    locale_code            String,
    continent_code         String,
    continent_name         String,
    country_iso_code       String,
    country_name           String,
    subdivision_1_iso_code String,
    subdivision_1_name     String,
    subdivision_2_iso_code String,
    subdivision_2_name     String,
    city_name              String,
    metro_code             UInt32,
    time_zone              String,
    is_in_european_union   UInt8
) ENGINE = Dictionary('geoip_city_locations_zh');

create table if not exists geoip_city_blocks_ipv4
(
    prefix                         String,
    geoname_id                     UInt32,
    registered_country_geoname_id  UInt32,
    represented_country_geoname_id UInt32,
    postal_code                    String,
    latitude                       Float32,
    longitude                      Float32,
    accuracy_radius                UInt32
) ENGINE = Dictionary('geoip_city_blocks_ipv4');

create table if not exists geoip_city_blocks_ipv6
(
    prefix                         String,
    geoname_id                     UInt32,
    registered_country_geoname_id  UInt32,
    represented_country_geoname_id UInt32,
    postal_code                    String,
    latitude                       Float32,
    longitude                      Float32,
    accuracy_radius                UInt32
) ENGINE = Dictionary('geoip_city_blocks_ipv6');

create table if not exists geoip_country_locations_de
(
    geoname_id           UInt64,
    locale_code          String,
    continent_code       String,
    continent_name       String,
    country_iso_code     String,
    country_name         String,
    is_in_european_union UInt8
) ENGINE = Dictionary('geoip_country_locations_de');

create table if not exists geoip_country_locations_en
(
    geoname_id           UInt64,
    locale_code          String,
    continent_code       String,
    continent_name       String,
    country_iso_code     String,
    country_name         String,
    is_in_european_union UInt8
) ENGINE = Dictionary('geoip_country_locations_en');

create table if not exists geoip_country_locations_es
(
    geoname_id           UInt64,
    locale_code          String,
    continent_code       String,
    continent_name       String,
    country_iso_code     String,
    country_name         String,
    is_in_european_union UInt8
) ENGINE = Dictionary('geoip_country_locations_es');

create table if not exists geoip_country_locations_fr
(
    geoname_id           UInt64,
    locale_code          String,
    continent_code       String,
    continent_name       String,
    country_iso_code     String,
    country_name         String,
    is_in_european_union UInt8
) ENGINE = Dictionary('geoip_country_locations_fr');

create table if not exists geoip_country_locations_ja
(
    geoname_id           UInt64,
    locale_code          String,
    continent_code       String,
    continent_name       String,
    country_iso_code     String,
    country_name         String,
    is_in_european_union UInt8
) ENGINE = Dictionary('geoip_country_locations_ja');

create table if not exists geoip_country_locations_pt
(
    geoname_id           UInt64,
    locale_code          String,
    continent_code       String,
    continent_name       String,
    country_iso_code     String,
    country_name         String,
    is_in_european_union UInt8
) ENGINE = Dictionary('geoip_country_locations_pt');

create table if not exists geoip_country_locations_ru
(
    geoname_id           UInt64,
    locale_code          String,
    continent_code       String,
    continent_name       String,
    country_iso_code     String,
    country_name         String,
    is_in_european_union UInt8
) ENGINE = Dictionary('geoip_country_locations_ru');

create table if not exists geoip_country_locations_zh
(
    geoname_id           UInt64,
    locale_code          String,
    continent_code       String,
    continent_name       String,
    country_iso_code     String,
    country_name         String,
    is_in_european_union UInt8
) ENGINE = Dictionary('geoip_country_locations_zh');

create table if not exists geoip_country_blocks_ipv4
(
    prefix                         String,
    geoname_id                     UInt32,
    registered_country_geoname_id  UInt32,
    represented_country_geoname_id UInt32
) ENGINE = Dictionary('geoip_country_blocks_ipv4');

create table if not exists geoip_country_blocks_ipv6
(
    prefix                         String,
    geoname_id                     UInt32,
    registered_country_geoname_id  UInt32,
    represented_country_geoname_id UInt32
) ENGINE = Dictionary('geoip_country_blocks_ipv6');
