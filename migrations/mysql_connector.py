import mysql.connector

# Read markdown content from file
with open("tutorial.md", "r") as file:
    markdown_content = file.read()

# Connect to your database
conn = mysql.connector.connect(
    host="localhost",
    user="your_username",
    password="your_password",
    database="your_database"
)
cursor = conn.cursor()

# Insert blog post
query = """
    INSERT INTO BlogPosts (title, author, content, tags, subject)
    VALUES (%s, %s, %s, %s, %s)
"""
values = (
    "Getting Started with FastAPI",
    "Keith Thomson",
    markdown_content,
    "fastapi,python,webdev",
    "Web Development"
)

cursor.execute(query, values)
conn.commit()

cursor.close()
conn.close()
