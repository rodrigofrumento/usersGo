CREATE TABLE address (
  id VARCHAR(36) NOT NULL PRIMARY KEY,
  cep VARCHAR(255) NOT NULL,
  ibge VARCHAR(255) NOT NULL,
  uf VARCHAR(255) NOT NULL,
  city VARCHAR(255) NOT NULL,
  complement VARCHAR(255) NULL,
  street VARCHAR(255) NOT NULL,
  created_at TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP(3) NOT NULL,
  user_id VARCHAR(36) UNIQUE NOT NULL,
  
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);