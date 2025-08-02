# HyperTube
Modern video streaming web app with real-time torrent downloading. Built with Vue.js frontend and Go backend. Features movie search across multiple APIs, progressive streaming while downloading, format conversion, and multi-language subtitle support.


hypertube/
├── frontend/                 # Vue.js application
│   ├── src/
│   │   ├── components/
│   │   │   ├── auth/
│   │   │   ├── movie/
│   │   │   ├── user/
│   │   │   └── common/
│   │   ├── views/
│   │   ├── router/
│   │   ├── store/           # Vuex/Pinia
│   │   ├── services/        # API calls
│   │   ├── utils/
│   │   └── assets/
│   ├── public/
│   └── package.json
│
├── backend/                 # Go application
│   ├── cmd/
│   │   └── server/
│   │       └── main.go
│   ├── internal/
│   │   ├── auth/           # OAuth, JWT handling
│   │   ├── torrent/        # BitTorrent client
│   │   ├── video/          # Video processing, streaming
│   │   ├── api/            # REST API handlers
│   │   ├── models/         # Database models
│   │   ├── services/       # Business logic
│   │   ├── middleware/     # Authentication, CORS, etc.
│   │   └── database/       # DB connection, migrations
│   ├── pkg/                # Shared packages
│   ├── scripts/            # Database migrations
│   ├── storage/            # Downloaded movies, temp files
│   └── go.mod
│
├── docker-compose.yml      # PostgreSQL, Redis
├── .env.example
└── README.md