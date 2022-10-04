# Go Admin App
A simple backend app/API that lets you create/manage users, products and their orders. Source code is based on [this Udemy course](https://udemy.com/course/the-complete-react-golang-course). The server runs on port 3000 and depends on a MySQL database.

# Code Structure
## [controllers](./controllers)
This directory contains the codes that handle all operations on specific database models

### [authController.go](./controllers/authController.go)
<table>
<tr>
    <th>Function</th>
    <th class="description">Description</th>
    <th>API Endpoint</th>
    <th>Method</th>
    <th>Body/Params</th>
</tr>

<!-- Data Begins -->
<tr>
    <td class="function">Register</td>
    <td class="description">Registers a new user</td>
    <td class="api-endpoint"><code>/api/register</code></td>
    <td class="method"><code>POST</code></td>
    <td class="body-params">
    <pre class="body-params">
    json
    {
        "first_name": "Harry",
        "last_name": "Potter",
        "email": "hjp@hogwarts.edu",
        "password": "patronus",
        "password_confirm": "patronus"
    }
    </pre>
    </td>
</tr>
<tr>
    <td class="function">Login</td>
    <td class="description">
    Logs in a user</br>
    Set a JWT token in the cookie.
    </td>
    <td class="api-endpoint"><code>/api/login</code></td>
    <td class="method"><code>POST</code></td>
    <td class="body-params">
    <pre class="body-params">
    json
    {
        "email": "hjp@hogwarts.edu",
        "password": "patronus",
    }
    </pre>
    </td>
</tr>
<tr>
    <td class="function">User</td>
    <td class="description">Get the logged-in user's info (except password)</td>
    <td class="api-endpoint"><code>/api/user</code></td>
    <td class="method"><code>GET</code></td>
    <td class="body-params">
    <pre></pre>
    </td>
</tr>
<tr>
    <td class="function">Logout</td>
    <td class="description">
    Logs out the logged-in user</br>
    Removes the JWT token cookie.
    </td>
    <td class="api-endpoint"><code>/api/logout</code></td>
    <td class="method"><code>POST</code></td>
    <td class="body-params">
    <pre></pre>
    </td>
</tr>
<tr>
    <td class="function">UpdateInfo</td>
    <td class="description">Updates the logged-in user's info (except password)</td>
    <td class="api-endpoint"><code>/api/users/info</code></td>
    <td class="method"><code>PUT</code></td>
    <td class="body-params">
    <pre class="body-params">
    json
    {
        "first_name": "Harry",
        "last_name": "Potter",
        "email": "hjp@hogwarts.edu",
    }
    </pre>
    </td>
</tr>
<tr>
    <td class="function">UpdatePassword</td>
    <td class="description">Updates the password for the logged-in user</td>
    <td class="api-endpoint"><code>/api/users/password</code></td>
    <td class="method"><code>PUT</code></td>
    <td class="body-params">
    <pre class="body-params">
    json
    {
        "password": "crucio",
        "password_confirm": "crucio"
    }
    </pre>
    </td>
</tr>
<!-- Data Ends -->
</table>

### [imageController](./controllers/imageController.go)
<table>
<tr>
    <th>Function</th>
    <th class="description">Description</th>
    <th>API Endpoint</th>
    <th>Method</th>
    <th>Body/Params</th>
</tr>

<!-- Data Begins -->
<tr>
    <td class="function">Upload</td>
    <td class="description">Uploads an image to <code>uploads</code> directory</td>
    <td class="api-endpoint"><code>/api/upload</code></td>
    <td class="method"><code>POST</code></td>
    <td class="body-params">
    <pre class="body-params">
    form-data
    {
        "image": "file_blob"
    }
    </pre>
    </td>
</tr>
<!-- Data Ends -->
</table>

### [orderController](./controllers/orderController.go)
<table>
<tr>
    <th>Function</th>
    <th class="description">Description</th>
    <th>API Endpoint</th>
    <th>Method</th>
    <th>Body/Params</th>
</tr>

<!-- Data Begins -->
<tr>
    <td class="function">AllOrders</td>
    <td class="description">Gets all orders in the DB, and paginates the result</td>
    <td class="api-endpoint"><code>/api/orders</code></td>
    <td class="method"><code>GET</code></td>
    <td class="body-params">
    <pre></pre>
    </td>
</tr>
<tr>
    <td class="function">Export</td>
    <td class="description">Exports orders into a CSV file in <code>csv/orders.csv</code></td>
    <td class="api-endpoint"><code>/api/export</code></td>
    <td class="method"><code>POST</code></td>
    <td class="body-params">
    <pre></pre>
    </td>
</tr>
<tr>
    <td class="function">Chart</td>
    <td class="description">Returns per-day orders</td>
    <td class="api-endpoint"><code>/api/chart</code></td>
    <td class="method"><code>GET</code></td>
    <td class="body-params">
    <pre></pre>
    </td>
</tr>
<!-- Data Ends -->
</table>

### [permissionController](./controllers/permissionController.go)

<table>
<tr>
    <th>Function</th>
    <th class="description">Description</th>
    <th>API Endpoint</th>
    <th>Method</th>
    <th>Body/Params</th>
</tr>

<!-- Data Begins -->
<tr>
    <td class="function">AllPermissions</td>
    <td class="description">Get all permission types from DB</td>
    <td class="api-endpoint"><code>/api/permissions</code></td>
    <td class="method"><code>GET</code></td>
    <td class="body-params">
    <pre></pre>
    </td>
</tr>
<!-- Data Ends -->
</table>

### [productController](./controllers/productController.go)

<table>
<tr>
    <th>Function</th>
    <th class="description">Description</th>
    <th>API Endpoint</th>
    <th>Method</th>
    <th>Body/Params</th>
</tr>

<!-- Data Begins -->
<tr>
    <td class="function">AllProducts</td>
    <td class="description">Get all the products from DB, and paginate result</td>
    <td class="api-endpoint"><code>/api/products</code></td>
    <td class="method"><code>GET</code></td>
    <td class="body-params">
    <pre></pre>
    </td>
</tr>
<tr>
    <td class="function">CreateProduct</td>
    <td class="description">Create a Product in DB.</td>
    <td class="api-endpoint"><code>/api/products</code></td>
    <td class="method"><code>POST</code></td>
    <td class="body-params">
    <pre class="body-params">
    json
    {
        "title": "Wand",
        "description": "MX1000",
        "image": "not found",
        "price": 1000
    }
    </pre>
    </td>
</tr>
<tr>
    <td class="function">GetProduct</td>
    <td class="description">Get the product with specified id</td>
    <td class="api-endpoint"><code>/api/products/{id}</code></td>
    <td class="method"><code>GET</code></td>
    <td class="body-params">
    <pre></pre>
    </td>
</tr>
<tr>
    <td class="function">UpdateProduct</td>
    <td class="description">Update the product details of a specific product</td>
    <td class="api-endpoint"><code>/api/products/{id}</code></td>
    <td class="method"><code>PUT</code></td>
    <td class="body-params">
    <pre class="body-params">
    json
    {
        "title": "Wand",
        "description": "MX1000",
        "image": "not found",
        "price": 1100
    }
    </pre>
    </td>
</tr>
<tr>
    <td class="function">DeleteProduct</td>
    <td class="description">Deletes the specified product</td>
    <td class="api-endpoint"><code>/api/products/{id}</code></td>
    <td class="method"><code>DELETE</code></td>
    <td class="body-params">
    <pre></pre>
    </td>
</tr>
<!-- Data Ends -->
</table>

### [roleController](./controllers/roleController.go)

<table>
<tr>
    <th>Function</th>
    <th class="description">Description</th>
    <th>API Endpoint</th>
    <th>Method</th>
    <th>Body/Params</th>
</tr>

<!-- Data Begins -->
<tr>
    <td class="function">AllRoles</td>
    <td class="description">Get all the roles from DB, and paginate result</td>
    <td class="api-endpoint"><code>/api/roles</code></td>
    <td class="method"><code>GET</code></td>
    <td class="body-params">
    <pre></pre>
    </td>
</tr>
<tr>
    <td class="function">CreateRole</td>
    <td class="description">Create a Role in DB.</td>
    <td class="api-endpoint"><code>/api/roles</code></td>
    <td class="method"><code>POST</code></td>
    <td class="body-params">
    <pre class="body-params">
    json
    {
        "name": "Administrator"
    }
    </pre>
    </td>
</tr>
<tr>
    <td class="function">GetRole</td>
    <td class="description">Get the role with specified id</td>
    <td class="api-endpoint"><code>/api/roles/{id}</code></td>
    <td class="method"><code>GET</code></td>
    <td class="body-params">
    <pre></pre>
    </td>
</tr>
<tr>
    <td class="function">UpdateRole</td>
    <td class="description">Update the role details of a specific role</td>
    <td class="api-endpoint"><code>/api/roles/{id}</code></td>
    <td class="method"><code>PUT</code></td>
    <td class="body-params">
    <pre class="body-params">
    json
    {
        "name": "Admin",
    }
    </pre>
    </td>
</tr>
<tr>
    <td class="function">DeleteRole</td>
    <td class="description">Deletes the specified role</td>
    <td class="api-endpoint"><code>/api/roles/{id}</code></td>
    <td class="method"><code>DELETE</code></td>
    <td class="body-params">
    <pre></pre>
    </td>
</tr>
<!-- Data Ends -->
</table>

### [userController](./controllers/userController.go)

<table>
<tr>
    <th>Function</th>
    <th class="description">Description</th>
    <th>API Endpoint</th>
    <th>Method</th>
    <th>Body/Params</th>
</tr>

<!-- Data Begins -->
<tr>
    <td class="function">AllUsers</td>
    <td class="description">Get all the users from DB, and paginate result</td>
    <td class="api-endpoint"><code>/api/users</code></td>
    <td class="method"><code>GET</code></td>
    <td class="body-params">
    <pre></pre>
    </td>
</tr>
<tr>
    <td class="function">CreateUser</td>
    <td class="description">Create a User in DB.</td>
    <td class="api-endpoint"><code>/api/users</code></td>
    <td class="method"><code>POST</code></td>
    <td class="body-params">
    <pre class="body-params">
    json
    {
        "first_name": "Hermione",
        "last_name": "Granger"
        "email": "hjg@hogwarts.edu"
    }
    </pre>
    </td>
</tr>
<tr>
    <td class="function">GetUser</td>
    <td class="description">Get the user with specified id</td>
    <td class="api-endpoint"><code>/api/users/{id}</code></td>
    <td class="method"><code>GET</code></td>
    <td class="body-params">
    <pre></pre>
    </td>
</tr>
<tr>
    <td class="function">UpdateUser</td>
    <td class="description">Update the user details of a specific user</td>
    <td class="api-endpoint"><code>/api/users/{id}</code></td>
    <td class="method"><code>PUT</code></td>
    <td class="body-params">
    <pre class="body-params">
    json
    {
        "first_name": "Hermione",
        "last_name": "Granger"
        "email": "hjg@hogwarts.edu"
    }
    </pre>
    </td>
</tr>
<tr>
    <td class="function">DeleteUser</td>
    <td class="description">Deletes the specified user</td>
    <td class="api-endpoint"><code>/api/users/{id}</code></td>
    <td class="method"><code>DELETE</code></td>
    <td class="body-params">
    <pre></pre>
    </td>
</tr>
<!-- Data Ends -->
</table>

