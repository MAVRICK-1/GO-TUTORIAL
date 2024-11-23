import pandas as pd
import json

# Load the CSV file
file_path = 'project.csv'
df = pd.read_csv(file_path)

# Filter for blockchain-based projects
blockchain_projects = df[df['Is your Project is based on Blockchain?'].str.contains('yes', case=False, na=False)]

# Create a list of dictionaries with the required fields
blockchain_projects_data = blockchain_projects.apply(lambda row: {
    "name": row['Name of the Project Owner'],
    "title": row['Name of the project '],
    "tags": [tag.strip() for tag in row['Tech Stacks '].split(',')] if pd.notna(row['Tech Stacks ']) else [],
    "description": row['Description of the Project (for socials and website) '] if pd.notna(row['Description of the Project (for socials and website) ']) else "",
    "projectUrl": row['Complete Repository URL '] if pd.notna(row['Complete Repository URL ']) else ""
}, axis=1).tolist()

# Convert the list of blockchain projects into JSON format
blockchain_projects_json = json.dumps(blockchain_projects_data, indent=4)
output_file = 'project.json'
# Write the JSON structure to a file
with open(output_file, 'w') as f:
    f.write(blockchain_projects_json)

print(f'JSON data successfully written to {output_file}') #theis 
