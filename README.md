# LogGuard

LogGuard is a lightweight log monitoring security tool built with Go and TypeScript.

It detects suspicious IP activity based on access frequency and provides results via a web interface.

---

## Features

- REST API (Go)
- Frontend UI (Vite + TypeScript)
- CORS-enabled API integration
- Lightweight and easy to run

---

## Architecture

Sample Log → Go API (IP parsing & counting) → JSON response → TypeScript Frontend (display)

This project demonstrates a simple Intrusion Detection System (IDS) concept by analyzing access frequency patterns, simulating basic security monitoring behavior.

---

## Usage

### Backend

```bash
go run main.go
```

### Frontend

```bash
cd frontend
npm install
npm run dev
```

Access:  
http://localhost:5173

---

## Example Output

```
⚠ Suspicious IP detected: 192.168.1.1 (access count: 12)
⚠ Suspicious IP detected: 10.0.0.1 (access count: 7)
```

---

## Future Improvements

- Log file parsing (real data)
- CSV report export
- Hash integrity check
- Severity classification
