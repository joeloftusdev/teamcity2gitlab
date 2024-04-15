# GitLab and TeamCity Setup Script

This script automates the process of setting up GitLab, a GitLab runner, and TeamCity using Docker.

## Prerequisites

- Docker and Docker Compose installed on your machine.
- The script file `setup.sh` is in your current directory /lab/.

## Steps to Run the Script

1. Open a terminal.

2. Navigate to the directory where you saved `setup.sh`.

3. Make the script executable with the following command:

    ```bash
    chmod +x setup.sh
    ```

4. Run the script:

    ```bash
    ./setup.sh
    ```

5. The script will start GitLab and a GitLab runner using Docker Compose. It will then wait for GitLab to start.

6. Once GitLab has started, the script will retrieve the root password for GitLab and print it. It will also print the URL for GitLab.

7. Open a web browser and navigate to the printed URL. Log in to GitLab using the username `root` and the printed password.

8. Register a new runner in GitLab and obtain the registration token.

9. Go back to the terminal where the script is running. Enter the registration token when prompted.

10. The script will register the runner and update the runner's configuration.

11. The script will then navigate to the TeamCity directory and start TeamCity using Docker Compose.

12. The script will wait for TeamCity to start and then print the URL for TeamCity.

13. Open a web browser and navigate to the printed URL. Follow the TeamCity start up wizard.

14. Accept the license agreement then go back to the terminal and press any key to continue.

15. The script will retrieve the super user authentication token for TeamCity and print it.

16. Log in to TeamCity using the printed token as the password (the username is blank).

17. Once the script has finished running, your GitLab and TeamCity setup is complete.

## Note

This script assumes that GitLab will be accessible at `http://localhost:8080` and TeamCity at `http://localhost:8111`. If you're running these services at different URLs, you will need to modify the script.