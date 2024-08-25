# CRM System

## Overview
This project is a CRM backend system built using Go and the Gin framework. It includes user authentication, customer and lead management, interaction tracking, and more.

## Features
- User Management (CRUD operations)
- Interaction Tracking
- Advanced Analytics (Optional)
- Email Integration (Optional)
- Activity Notifications

## Database Schema
**Users**
- ID
- Name
- Contact Information
- Company
- Status
- Notes

**Customers**
- ID
- Name
- Contact Information
- Status
- Notes

**Interactions**
- ID
- Customer ID
- Interaction Type
- Status
- Details

## Setup
1. Clone the repository:
   git clone https://github.com/your-repo/crm-system.git
   cd crm-system

2. Build and run the Docker containers:
    docker-compose up --build

3. Access the API:
    Open http://localhost:8080 in browser.

Deployment
Follow `deploy/aws-deploy.sh` for deployment instructions.
