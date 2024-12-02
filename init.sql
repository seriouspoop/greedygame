CREATE schema IF NOT EXISTS delivery;

CREATE TABLE delivery.targeting_rule_rec (
    campaign_id TEXT NOT NULL,         
    app_include TEXT[],                
    app_exclude TEXT[],                
    country_include TEXT[],            
    country_exclude TEXT[],            
    os_include TEXT[],                 
    os_exclude TEXT[],                 
    PRIMARY KEY (campaign_id)          
);

CREATE TABLE delivery.campaign (
    campaign_id TEXT NOT NULL,
    name TEXT NOT NULL,
    image TEXT NOT NULL,
    cta TEXT NOT NULL,
    status TEXT NOT NULL,
    PRIMARY KEY (campaign_id)
);