# distributed-compression-storage

# TODO List 

- [ ] **Add Metadata Management**
  - [ ] Store metadata (e.g., file name, original size, compression ratio, chunk locations).
  - [ ] Use a database like SQLite or PostgreSQL for persistence.

- [ ] **Add Retrieval Functionality**
  - [ ] Implement a function to fetch and reassemble chunks into the original file.
  - [ ] Test the upload-download cycle.

- [ ] **Support Multiple Compression Algorithms**
  - [ ] Add support for GZIP, Zstandard, and Snappy.
  - [ ] Allow users to select compression algorithms.

---

- [ ] **Implement Microservices**
  - [ ] Split into services: Chunking, Compression, Storage, Orchestrator.
  - [ ] Define APIs for communication.

- [ ] **Use gRPC for Communication**
  - [ ] Replace function calls with gRPC between services.

- [ ] **Run on Kubernetes**
  - [ ] Use Helm charts for configuration.
  - [ ] Deploy services and test scalability.

---

- [ ] **Integrate Distributed Storage**
  - [ ] Store chunks across AWS S3, Google Cloud Storage, or Azure Blob Storage.
  - [ ] Implement redundancy and chunk distribution logic.

- [ ] **Improve Error Handling & Resilience**
  - [ ] Add retry mechanisms for failed uploads/downloads.
  - [ ] Handle network failures and storage degradation.

- [ ] **Large File Support**
  - [ ] Optimize handling for multi-GB files using streaming APIs.

---

- [ ] **Develop a Command-Line Interface (CLI)**
  - [ ] Create commands for uploading, downloading, and checking file status.

- [ ] **Build a Web Interface**
  - [ ] Use Go’s `net/http` or React.js to provide an upload/download UI.

---

- [ ] **Write Unit and Integration Tests**
  - [ ] Test all functions and interactions.
  - [ ] Use mocks for external dependencies.

- [ ] **Benchmark Performance**
  - [ ] Test with various file sizes and types.
  - [ ] Identify and resolve bottlenecks.

- [ ] **Handle Edge Cases**
  - [ ] Test empty files, special file names, and corrupted chunks.

---

- [ ] **Update `README.md`**
  - [ ] Add detailed project description, installation, and usage instructions.
  - [ ] Include examples and screenshots.

- [ ] **Document APIs**
  - [ ] Write clear documentation for microservice APIs.

- [ ] **Create Architecture Diagram**
  - [ ] Visualize system components and their interactions.

---

- [ ] **Dockerize the System**
  - [ ] Write a `Dockerfile` for each service.
  - [ ] Create a `docker-compose.yml` for local testing.

- [ ] **Deploy a Demo**
  - [ ] Deploy the system to a cloud platform (e.g., AWS, GCP, or DigitalOcean).
  - [ ] Provide a public endpoint or demo UI.

---

## Maybe...
- [ ] **Add File Encryption**
  - [ ] Encrypt chunks before storage for security.

- [ ] **Implement Deduplication**
  - [ ] Use hashing to avoid storing duplicate chunks.

- [ ] **Ensure Data Integrity**
  - [ ] Add checksums or hashes for chunk validation.

- [ ] **Add Analytics**
  - [ ] Track metrics like compression ratio, storage usage, and access frequency.