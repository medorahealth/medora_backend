CREATE TABLE Users (
    userID INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    phone_number VARCHAR(15) NOT NULL UNIQUE,
    whatsapp_number VARCHAR(15) NOT NULL,
    email VARCHAR(254) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    razorpay_customer_id VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    refresh_token TEXT,
    is_active BOOLEAN NOT NULL DEFAULT TRUE
);

CREATE TABLE Labs (
    labID BIGSERIAL PRIMARY KEY,
    operating_dr_id BIGINT REFERENCES users(id) ON DELETE RESTRICT,
    legal_name TEXT NOT NULL,
    description TEXT,
    accreditation TEXT NOT NULL,
    license_number TEXT,
    license_pdf TEXT NOT NULL,
    pollution_policy_pdf TEXT NOT NULL,
    contact_email TEXT,
    contact_phone TEXT,

    -- Address and geo
    line1 TEXT NOT NULL,
    line2 TEXT,
    city TEXT NOT NULL,
    state TEXT NOT NULL,
    pincode TEXT NOT NULL,
    latitude DOUBLE PRECISION,
    longitude DOUBLE PRECISION,

    -- Simple operating hours and service flags
    opens_at TIME,
    closes_at TIME,
    emergency_service BOOLEAN NOT NULL DEFAULT FALSE,
    
    -- Status and ratings
    is_active BOOLEAN NOT NULL DEFAULT TRUE,
    rating NUMERIC(3,2) DEFAULT 0.0,
    ratings_count INT DEFAULT 0,

    created_at TIMESTAMP NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMP NOT NULL DEFAULT NOW(),

    -- The CHECK constraint to validate the state name against the official list
    CONSTRAINT check_indian_state CHECK (state IN (
        'Andhra Pradesh',
        'Arunachal Pradesh',
        'Assam',
        'Bihar',
        'Chhattisgarh',
        'Goa',
        'Gujarat',
        'Haryana',
        'Himachal Pradesh',
        'Jharkhand',
        'Karnataka',
        'Kerala',
        'Madhya Pradesh',
        'Maharashtra',
        'Manipur',
        'Meghalaya',
        'Mizoram',
        'Nagaland',
        'Odisha',
        'Punjab',
        'Rajasthan',
        'Sikkim',
        'Tamil Nadu',
        'Telangana',
        'Tripura',
        'Uttar Pradesh',
        'Uttarakhand',
        'West Bengal',
        'Andaman and Nicobar Islands',
        'Chandigarh',
        'Dadra and Nagar Haveli and Daman and Diu',
        'Delhi',
        'Jammu and Kashmir',
        'Ladakh',
        'Lakshadweep',
        'Puducherry'
    ))
);
