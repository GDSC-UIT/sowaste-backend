<style>
    img {
        border-radius: 10px;
        border: 2px dashed;
    }
</style>

# **/internal**

The heart of our application, i.e., all its internal logic, is stored here. `/internal` is not imported into other applications and libraries. The code written here is intended solely for internal use within the code base. Starting from the `Go 1.4 version`, a defined mechanism has been in place that prevents the importing of packages outside of this project if they are inside `/internal`.

`/internal` is where we store the business logic of the project, together with all work related to databases. In other words, all the logic associated with this app. The structure inside `/internal` can be organized in a variety of ways, depending on the specific architecture used. I’m not going to go into this in too much depth right now, but I will illustrate in broad strokes what it looks like. Here is an example of a three-layer architecture, where the app is divided into three layers:

- Transport.
- Business.
- Databases.

The logic should be such that the layers hierarchically communicate with each other from top to bottom, and vice versa. No layer may “skip” over its intermediate peer (e.g. when the transport layer communicates directly with the database) and no layer that is below another one may communicate directly with the layer above (e.g. when the database communicates with the transport layer).

<img src="https://miro.medium.com/max/828/0*hX3AzeOYrYQknRpO" alt="A three-tiered architecture model" />

## **The transport layer:**

The network layer of the application, where the end user interacts with the app. Once the request has been processed, all the information collected goes to the layer below.

## **The business layer:**

As the name implies, this layer contains the business logic that supports the app’s core functions. If the logic involves databases, we move down to the layer below.

## **The database layer:**

This layer is responsible for interacting with permanent vaults, such as databases, and other non-business-related information processing. For instance, reading and writing in the database.

## **/internal directories:**

#### `/config`

Initialization of the general app configurations that we wrote in the root of the project.

#### `/database (the database layer)`

The files contain methods for interacting with databases.

#### `/models (the database layer)`

The structures of database tables.

#### `/services (the business layer)`

The entire business logic of the application.

#### `/transport (the transport layer)`

Here we store http-server settings, handlers, ports, etc.

<img src="https://miro.medium.com/max/828/0*IoQPbZsLkWwbjJdd" alt="Structure of the /internal directory" />
