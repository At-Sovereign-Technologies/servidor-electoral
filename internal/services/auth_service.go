package services

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"encoding/pem"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"time"
)

type AuthService struct {
	configPath string
	caCert     *x509.Certificate
	caKey      *rsa.PrivateKey
}

func NewAuthService(configPath string) (*AuthService, error) {
	authService := &AuthService{configPath: configPath}

	if err := authService.loadOrGenerateCertificate(); err != nil {
		return nil, fmt.Errorf("failed to load or generate certificate: %w", err)
	}

	return authService, nil
}

func (s *AuthService) generateCertificate() error {
	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return err
	}

	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject: pkix.Name{
			Organization: []string{"Sello Legitimo"},
			CommonName:   "sello-legitimo-ca",
		},
		IsCA:                  true,
		BasicConstraintsValid: true,
		NotBefore:             time.Now(),
		NotAfter:              time.Now().Add(5 * 365 * 24 * time.Hour),
		KeyUsage:              x509.KeyUsageCertSign | x509.KeyUsageDigitalSignature,
	}

	certBytes, err := x509.CreateCertificate(rand.Reader, template, template, &privateKey.PublicKey, privateKey)
	if err != nil {
		return err
	}

	certPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})

	keyPEM := pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	if err := os.WriteFile(filepath.Join(s.configPath, "ca.crt"), certPEM, 0644); err != nil {
		return err
	}

	if err := os.WriteFile(filepath.Join(s.configPath, "ca.key"), keyPEM, 0600); err != nil {
		return err
	}

	cert, _ := x509.ParseCertificate(certBytes)

	s.caCert = cert
	s.caKey = privateKey

	return nil
}

func (s *AuthService) loadCertificate() error {
	certPEM, err := os.ReadFile(filepath.Join(s.configPath, "ca.crt"))
	if err != nil {
		return err
	}

	keyPEM, err := os.ReadFile(filepath.Join(s.configPath, "ca.key"))
	if err != nil {
		return err
	}

	certBlock, _ := pem.Decode(certPEM)
	keyBlock, _ := pem.Decode(keyPEM)

	if certBlock == nil || keyBlock == nil {
		return fmt.Errorf("failed to decode PEM")
	}

	cert, err := x509.ParseCertificate(certBlock.Bytes)
	if err != nil {
		return err
	}

	key, err := x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		return err
	}

	s.caCert = cert
	s.caKey = key

	return nil
}

func (s *AuthService) loadOrGenerateCertificate() error {
	if err := s.loadCertificate(); err == nil {
		return nil
	}

	return s.generateCertificate()
}

func (s *AuthService) IssueCertificate(commonName string) (certPEM, keyPEM []byte, err error) {
	if err := s.loadOrGenerateCertificate(); err != nil {
		return nil, nil, err
	}

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return nil, nil, err
	}

	serial, _ := rand.Int(rand.Reader, big.NewInt(1<<62))

	template := &x509.Certificate{
		SerialNumber: serial,
		Subject: pkix.Name{
			CommonName: commonName,
		},
		NotBefore: time.Now(),
		NotAfter:  time.Now().Add(24 * time.Hour), // short-lived

		KeyUsage: x509.KeyUsageDigitalSignature,
		ExtKeyUsage: []x509.ExtKeyUsage{
			x509.ExtKeyUsageClientAuth,
		},
	}

	certBytes, err := x509.CreateCertificate(
		rand.Reader,
		template,
		s.caCert,
		&privateKey.PublicKey,
		s.caKey,
	)
	if err != nil {
		return nil, nil, err
	}

	certPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "CERTIFICATE",
		Bytes: certBytes,
	})

	keyPEM = pem.EncodeToMemory(&pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: x509.MarshalPKCS1PrivateKey(privateKey),
	})

	return certPEM, keyPEM, nil
}

func (s *AuthService) GenerateSecret() (string, error) {
	randomBytes := make([]byte, 16)
	_, err := rand.Read(randomBytes)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", randomBytes), nil
}
