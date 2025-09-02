# Go Web Application CI/CD with Jenkins

This project demonstrates a **CI/CD pipeline** for a **Go web application** using **Jenkins**. It automates code fetching, building, and email notification on build status.

---

##  Project Overview
- Language: **Go**
- CI/CD Tool: **Jenkins**
- Repository: [GitHub - Go Web App](https://github.com/mikkilybhuvan/emailnotification)
- Features:
  - Automatic build triggers from GitHub commits
  - Email notifications on success/failure
  - Simple Go web application as an example

---

##  Project Structure
GoApp/
├── main.go # Simple Go web app
├── Jenkinsfile # Jenkins pipeline configuration (if using pipeline)
└── README.md # Documentation


---

##  Jenkins Job Setup
1. Install **Jenkins** and required plugins:
   - Git Plugin  
   - Email Extension Plugin  

2. Create a **Freestyle Job**:
   - Source Code Management → Git → Add Repository URL  
   - Branch → `main`  
   - Build → Execute shell:
     ```sh
     go build main.go
     ```
   - Post-build → Email Notification  

3. Run the job → Jenkins will fetch code, build, and notify via email.

---

##  Email Notifications
- Configured Jenkins to send email after every build.
- Example:  
  - ✅ Success → "Build Successful"  
  - ❌ Failure → "Build Failed"

---

##  How to Run Locally
```bash
# Clone repository
git clone https://github.com/mikkilybhuvan/emailnotification.git
cd emailnotification

# Run application
go run main.go

Key Learnings

How to create and configure a Jenkins Freestyle job.

Connecting GitHub repository with Jenkins.

Automating builds with Go applications.

Sending email notifications for CI/CD pipeline results.

Understanding the basic flow of Continuous Integration (CI).

 Conclusion

This project successfully shows a working CI/CD pipeline using Jenkins for a Go application.
It can be extended further with Docker, Kubernetes, or advanced pipeline scripting.
