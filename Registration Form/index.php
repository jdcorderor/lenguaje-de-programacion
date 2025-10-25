<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Registration Form</title>
    <link rel="stylesheet" href="styles.css">
</head>
<body>
    <div class="form-container">
        <h2>User Registration</h2>
        <form action="procesar.php" method="POST">
            <div class="form-group">
                <label for="nombre">Name *</label>
                <input type="text" id="nombre" name="nombre" required>
            </div>
            
            <div class="form-group">
                <label for="email">Email *</label>
                <input type="email" id="email" name="email" required>
            </div>
            
            <div class="form-group">
                <label for="password">Password *<span class="hint">(Must be at least 8 characters long, and include special characters)</span></label>
                <input type="password" id="password" name="password" required>
            </div>
            
            <button type="submit" class="submit-button">Register now!</button>
        </form>

        <p class="footer">Made by Eduardo Tovar and Juan Diego Cordero</p>
    </div>
</body>
</html>
