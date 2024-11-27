User consents (persetujuan pengguna) adalah komponen penting dalam OAuth2 yang berfungsi untuk mencatat dan mengelola izin yang diberikan pengguna kepada aplikasi client. Mari saya jelaskan lebih detail:
OAuth2 Consent Flow and ManagementClick to open diagram
Fungsi utama menyimpan user consents:

Transparansi dan Kontrol

Mencatat secara eksplisit izin apa saja yang diberikan user kepada aplikasi
User bisa melihat aplikasi mana saja yang memiliki akses ke akunnya
User bisa mencabut akses aplikasi kapan saja


Automated Authorization

Menghindari user harus menyetujui ulang setiap kali aplikasi yang sama meminta akses
Mempercepat proses autentikasi untuk aplikasi yang sudah disetujui sebelumnya


Granular Permission Control

Menyimpan scope spesifik yang diizinkan untuk setiap aplikasi
Memungkinkan user memberikan akses terbatas (misal: read-only)
Memungkinkan update izin tanpa harus mencabut semua akses



Contoh struktur tabel user_consents:
```aiignore
TABLE user_consents (
id UUID PRIMARY KEY,
user_id UUID NOT NULL,
client_id UUID NOT NULL,
scopes TEXT[] NOT NULL,
created_at TIMESTAMP NOT NULL,
last_used_at TIMESTAMP,
expires_at TIMESTAMP,
FOREIGN KEY (user_id) REFERENCES users(id),
FOREIGN KEY (client_id) REFERENCES clients(id)
);  
```

Use cases penting:

Security Audit

Melacak kapan dan kepada siapa user memberikan akses
Membantu dalam investigasi jika terjadi masalah keamanan
Memastikan compliance dengan regulasi privacy


User Experience
```curl
// Contoh response GET /oauth/consents
{
  "consents": [
    {
      "client_name": "Spotify App",
      "scopes": ["read_profile", "read_playlist"],
      "last_used": "2024-11-26T10:00:00Z",
      "granted_at": "2024-10-01T14:30:00Z"
    }
  ]
}
```
Revocation Management

Saat user mencabut akses:

Hapus consent record
Revoke semua active tokens
Batalkan semua refresh tokens
Kirim notifikasi ke client (optional)




Compliance & Privacy

Membantu memenuhi requirement GDPR tentang consent
Menyediakan audit trail untuk privacy compliance
Memungkinkan user mengexport data consent mereka



Best practices implementasi:

Selalu timestamp setiap consent record
Implement soft delete untuk history audit
Set up automatic cleanup untuk expired consents
Provide clear UI untuk user memanage consents
Implement webhook notifications untuk consent changes

Apakah Anda ingin saya menjelaskan lebih detail tentang implementasi spesifik atau aspek tertentu dari consent management?

```sequenceDiagram
participant U as User
participant C as Client App
participant AS as Auth Server
participant DB as Consent DB

    rect rgb(200, 220, 240)
        Note over U,DB: First Time Authorization Flow
        C->>AS: Request Authorization
        AS->>DB: Check Existing Consent
        DB->>AS: No Consent Found
        AS->>U: Display Consent Screen
        Note over U,AS: Shows requested permissions:<br/>- Read Profile<br/>- Access Email<br/>- etc
        U->>AS: Approve Permissions
        AS->>DB: Store Consent Record
        AS->>C: Grant Authorization
    end

    rect rgb(220, 240, 200)
        Note over U,DB: Subsequent Authorization
        C->>AS: Request Authorization
        AS->>DB: Check Existing Consent
        DB->>AS: Consent Found
        AS->>C: Auto-grant Authorization
        Note over C,AS: Skip consent screen<br/>if permissions match
    end

    rect rgb(240, 220, 200)
        Note over U,DB: Consent Management
        U->>AS: View Authorized Apps
        AS->>DB: Fetch User's Consents
        DB->>AS: Return Consent List
        AS->>U: Display Authorized Apps
        U->>AS: Revoke App Access
        AS->>DB: Delete Consent Record
        AS->>DB: Revoke Related Tokens
    end```


