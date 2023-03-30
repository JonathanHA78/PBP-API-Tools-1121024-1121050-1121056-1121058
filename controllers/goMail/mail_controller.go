package controllers

import (
	"apitools/controllers"
	"os"

	"gopkg.in/gomail.v2"
)

func sendEmail(templatePath string, content string, receiverMail string) string {
	m := gomail.NewMessage()
	m.SetHeader("From", "No Reply <no-reply@example.com>")
	m.SetHeader("To", receiverMail)
	m.SetHeader("Subject", "Hello!")
	m.SetBody("text/html", content)
	email := os.Getenv("EMAIL")
	password := os.Getenv("PASSWORD")
	d := gomail.NewDialer("smtp.gmail.com", 465, email, password)

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}
}

func generateEmail(emailType int, users []controllers.User) string {
	content := `<!DOCTYPE html>
	<html lang="en">
	<head>
		<meta charset="UTF-8">
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<style>
			body {
				font-family: Arial, sans-serif;
				font-size: 14px;
				line-height: 1.5;
				color: #333;
				margin: 0;
				padding: 0;
			}
	
			.email-container {
				max-width: 600px;
				margin: 0 auto;
				padding: 20px;
				background-color: #f8f8f8;
				border: 1px solid #ddd;
				border-radius: 5px;
			}
	
			.email-header {
				text-align: center;
				margin-bottom: 30px;
			}
	
			.email-header img {
				max-width: 150px;
				height: auto;
			}
	
			.email-content {
				padding: 20px;
				background-color: #ffffff;
				border-radius: 5px;
			}
	
			.email-content h1 {
				font-size: 24px;
				margin-bottom: 20px;
			}
	
			.email-content p {
				margin-bottom: 15px;
			}
	
			.email-content ul {
				padding-left: 20px;
				margin-bottom: 15px;
			}
	
			.email-footer {
				text-align: center;
				margin-top: 30px;
			}
	
			.email-footer p {
				font-size: 12px;
				color: #777;
			}
		</style>
	</head>
	<body>
		<div class="email-container">
			<div class="email-header">
				<img src="your-logo.png" alt="Your App Logo">
			</div>
			<div class="email-content">`
	switch emailType {
	case 1:
		//all of a user's tasklists
		content += `<h1>All Tasks</h1>
		<p>Hi ` + username + `,</p>
		<p>These are all of the tasks that you have noted: </p>
		`
		if tasks.length != nil {
			for i := 1; i <= tasks.length; i++ {
				content += `<ul><li><strong>Task Title:</strong> ` + tasks[i].title + `</li>
				<li><strong>Task Description:</strong> ` + tasks[i].description + `</li>
				<li><strong>Due Date:</strong> ` + tasks[i].due_date + `</li></ul>`
			}
			content += `
					<p>Please make sure to complete the task on time. If you have any questions or need assistance, feel free to contact us.</p>
					<p>Best regards,</p>
					<p>Your App Team</p></div>
					<div class="email-footer">
						<p>&copy; ` + current_year + `Your App. All Rights Reserved.</p>
					</div>
				</div>
			</body>
			</html>`
		} else {
			content += `<ul><li><strong>You have not created any tasks</strong> </li></ul>
					<p>Please explore our features. If you have any questions or need assistance, feel free to contact us.</p>
					<p>Best regards,</p>
					<p>Your App Team</p></div>
					<div class="email-footer">
						<p>&copy; ` + current_year + ` Your App. All Rights Reserved.</p>
					</div>
				</div>
			</body>
			</html>`
		}

	case 2:
		//user tasklist for the day
		content += `<h1>Tasks for the day</h1>
		<p>Hi ` + username + `,</p>
		<p>These are all of the tasks that you have set for today: </p>
		`
		if tasks.length != nil {
			for i := 1; i <= tasks.length; i++ {
				content += `<ul><li><strong>Task Title:</strong> ` + tasks[i].title + `</li>
				<li><strong>Task Description:</strong> ` + tasks[i].description + `</li>
				<li><strong>Due Date:</strong> ` + tasks[i].due_date + `</li></ul>`
			}
			content += `
					<p>Please make sure to complete the task on time. If you have any questions or need assistance, feel free to contact us.</p>
					<p>Best regards,</p>
					<p>Your App Team</p></div>
					<div class="email-footer">
						<p>&copy; ` + current_year + `Your App. All Rights Reserved.</p>
					</div>
				</div>
			</body>
			</html>`
		} else {
			content += `<ul><li><strong>You have no tasks for today</strong> </li></ul>
					<p>Please explore our features. If you have any questions or need assistance, feel free to contact us.</p>
					<p>Best regards,</p>
					<p>Your App Team</p></div>
					<div class="email-footer">
						<p>&copy; ` + current_year + ` Your App. All Rights Reserved.</p>
					</div>
				</div>
			</body>
			</html>`
		}
	case 3:
		//task creation
		content += `<h1>New Task Created</h1>
		<p>Hi ` + username + `,</p>
		<p>You have successfully created a new task in your to-do list: </p>
		`
		content += `<ul><li><strong>Task Title:</strong> ` + tasks[i].title + `</li>
			<li><strong>Task Description:</strong> ` + tasks[i].description + `</li>
			<li><strong>Due Date:</strong> ` + tasks[i].due_date + `</li></ul>
					<p>Keep up the good work and stay on top of your tasks. If you have any questions or need assistance, feel free to contact us.</p>
					<p>Best regards,</p>
					<p>Your App Team</p></div>
					<div class="email-footer">
						<p>&copy; ` + current_year + `Your App. All Rights Reserved.</p>
					</div>
				</div>
			</body>
			</html>`
	case 4:
		//task deletion
		content += `<h1>Task Deleted</h1>
		<p>Hi ` + username + `,</p>
		<p>We wanted to let you know that the following task has been successfully deleted from your to-do list:</p>
		`
		content += `<ul><li><strong>Task Title:</strong> ` + tasks[i].title + `</li>
			<li><strong>Task Description:</strong> ` + tasks[i].description + `</li>
			<li><strong>Due Date:</strong> ` + tasks[i].due_date + `</li></ul>
					<p>If you have deleted the task by mistake or need any assistance, feel free to contact us.</p>
					<p>Best regards,</p>
					<p>Your App Team</p></div>
					<div class="email-footer">
						<p>&copy; ` + current_year + `Your App. All Rights Reserved.</p>
					</div>
				</div>
			</body>
			</html>`
	case 5:
		//task update
		content += `<h1>Task Edited</h1>
		<p>Hi ` + username + `,</p>
		<p>You have successfully updated a new task in your to-do list: </p>
		`
		content += `<ul><li><strong>Task Title:</strong> ` + tasks[i].title + `</li>
			<li><strong>Task Description:</strong> ` + tasks[i].description + `</li>
			<li><strong>Due Date:</strong> ` + tasks[i].due_date + `</li></ul>
					<p>Keep up the good work and stay on top of your tasks. If you have any questions or need assistance, feel free to contact us.</p>
					<p>Best regards,</p>
					<p>Your App Team</p></div>
					<div class="email-footer">
						<p>&copy; ` + current_year + `Your App. All Rights Reserved.</p>
					</div>
				</div>
			</body>
			</html>`
	}

	return content
}
