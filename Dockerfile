# Use the official Node.js 18 image (LTS)
FROM node:18

# Create app directory inside container
WORKDIR /usr/src/app

# Copy package.json and package-lock.json (if any)
COPY package*.json ./

# Install app dependencies
RUN npm install --production

# Copy app source code
COPY . .

# Expose port 3000 for the app
EXPOSE 3000

# Run the app
CMD ["node", "index.js"]
