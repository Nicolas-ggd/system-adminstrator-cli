# System Administrator CLI
## ⚙️ Project Overview:
The goal is to create a
tool that automates system administration tasks, making life easier for sysadmins. This could be a command-line interface (CLI) application
that helps with monitoring, managing, and troubleshooting systems.

## 🌟 Project goals
Building a System Administrator Tool is an excellent way to apply Go skills to real-world problems. This project is a testing application to explore new possibilities and gain more knowledge in Go. Let's see what happens in the future!

## 🔨 Features:
1. System Monitoring: Monitor critical system metrics including CPU usage, memory consumption, memory swap, and network traffic.
2. Process Monitoring: View a detailed list of running processes, similar to the output of the top command in Linux.
3. Command-Line Interface: Simple and intuitive commands to fetch system information quickly.

# 🚀 Usage
This CLI tool offers three main commands:

1. `run`: Displays real-time statistics for CPU, memory, memory swap, and network usage.
    ```shell
    ./sysadmin-cli run
    ```
    ![Screenshot from 2024-08-12 19-08-46](https://github.com/user-attachments/assets/f61aedfb-04bf-4af6-b414-829b232a03c1)

2. `info`: Provides detailed information about the CPU, including statistics and usage.
    ```shell
    ./sysadmin-cli info
    ```
    ![Screenshot from 2024-08-12 19-07-52](https://github.com/user-attachments/assets/f4bb99f5-64cc-4ea4-adef-99ee5d0046d6)

3. `proc`: Lists all currently running processes, similar to the top command in Linux.
    ```shell
    ./sysadmin-cli info
    ```
    ![Screenshot from 2024-08-12 19-08-20](https://github.com/user-attachments/assets/893227f8-e7b4-49c1-94dc-5dc5e0147671)
    
## 🔨 Tools and Libraries:
1. Go's os/exec package: For running shell commands, executing scripts, or automating tasks.

## 🏗️ Contributing
Contributions are welcome! If anyone have interest to work with this project, hit a fork button....

## 📄 LICENSE
This project is licensed under the Apache-2.0 License.
