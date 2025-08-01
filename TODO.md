# TODO – Coral Project (Phase-based Roadmap)

## ✅ Phase 1 – Core Refactoring & Architecture
- [ ] Refactor OTP service to allow mockable interface for easier testing  
- [ ] Improve error handling structure across handlers  
- [ ] Add middleware for request logging and panic recovery  

## ✅ Phase 2 – API Documentation & Dev Tools
- [ ] Add Swagger/OpenAPI documentation  
- [ ] Write developer onboarding guide  

## ✅ Phase 3 – Testing Infrastructure
- [ ] Add integration tests for OTP flow  
- [ ] Increase test coverage for repository layer  
- [ ] Mock Kavenegar service in unit tests  

## ✅ Phase 4 – Deployment & CI/CD
- [ ] Dockerize `cmd/seed` with data population script  
- [ ] Create CI/CD pipeline (e.g., GitHub Actions or GitLab CI)  
- [ ] Add staging environment configuration  

## ✅ Phase 5 – Documentation for Users & Clients
- [ ] Add usage examples in README  
- [ ] Document error codes and API responses  

## ✅ Phase 6 – Functional Enhancements & Scaling
- [ ] Implement rate limiting for OTP requests  
- [ ] Role-based access control (RBAC)  
- [ ] Add Redis cache layer  
- [ ] Monitoring with Prometheus + Grafana  
- [ ] GraphQL support (optional)  
