# Use the official Node.js 16 image as the base
FROM node:latest

# Set the working directory inside the container
WORKDIR /app

# Copy package files to leverage caching
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the entire frontend directory
COPY . .

# Build the application (specific to Next.js; adjust if using another framework)
RUN npm run build

# Expose port 3000 for the frontend server
EXPOSE 3000

# Start the application
CMD ["npm", "start"]