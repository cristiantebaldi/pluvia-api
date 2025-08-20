CREATE TABLE usuarios (
    id SERIAL PRIMARY KEY,
    numero_telefone VARCHAR(25) NOT NULL,
    sms INT DEFAULT 0,
    whatsapp INT DEFAULT 0,
    criado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    atualizado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    habilitado INT DEFAULT 0
);

CREATE TABLE administradores (
    id SERIAL PRIMARY KEY,
    usuario VARCHAR(100) NOT NULL,
    senha VARCHAR(100) NOT NULL,
    email VARCHAR(100) NOT NULL,
    numero_telefone VARCHAR(25) NOT NULL,
    acesso INT DEFAULT 0,
    criado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    atualizado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE dispositivos (
    id SERIAL PRIMARY KEY,
    endereco_ip VARCHAR(50),
    endereco_mac VARCHAR(50)
);

CREATE TABLE dispositivos_administradores (
    id SERIAL PRIMARY KEY,
    usuario_id INT,
    dispositivo_id INT,
    acesso INT DEFAULT 0,
    criado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    atualizado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    CONSTRAINT fk_usuario FOREIGN KEY (usuario_id)
        REFERENCES administradores (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE,
    CONSTRAINT fk_dispositivo FOREIGN KEY (dispositivo_id)
        REFERENCES dispositivos (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE TABLE codigos (
    id SERIAL PRIMARY KEY,
    destinatario VARCHAR(100),
    codigo VARCHAR(100),
    criado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW()
);

CREATE TABLE notificacoes (
    id SERIAL PRIMARY KEY,
    dispositivo_administrador_id INT,
    criado_em TIMESTAMP WITH TIME ZONE DEFAULT NOW(),
    titulo VARCHAR(100) NOT NULL,
    texto TEXT NOT NULL,
    CONSTRAINT fk_dispositivo_administrador FOREIGN KEY (dispositivo_administrador_id)
        REFERENCES dispositivos_administradores (id)
        ON DELETE CASCADE
        ON UPDATE CASCADE
);

CREATE OR REPLACE FUNCTION atualizar_timestamp()
RETURNS TRIGGER AS $$
BEGIN
   NEW.atualizado_em = NOW();
   RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER trigger_atualizar_usuario
BEFORE UPDATE ON usuarios
FOR EACH ROW
EXECUTE FUNCTION atualizar_timestamp();

CREATE TRIGGER trigger_atualizar_administradores
BEFORE UPDATE ON administradores
FOR EACH ROW
EXECUTE FUNCTION atualizar_timestamp();

CREATE TRIGGER trigger_atualizar_dispositivos_administradores
BEFORE UPDATE ON dispositivos_administradores
FOR EACH ROW
EXECUTE FUNCTION atualizar_timestamp();