# React

## Installation
#### Create React App
  ```bash
  npm install -g create-react-app
  create-react-app <your app name>
  ```

## JSX
#### Concepts 
- Syntax extension to JavaScript.
- JSX produces React “elements”.

#### Differences between JSX and HTML
- JSX use “ClassName” rather than “Class” in HTML to define a class name.
- JSX use the camel-case naming convention.

#### Example
- Integrate JSX with JavaScript
  ```js
  const btnElement = <ahref="#"className="btn btn-primary">Buy</a>;
  ```
- Embed JavaScript code inside JSX
  ```js
  const btnName = "Buy";
  const btnClass = "btn btn-primary"; 
  const btnElement = <ahref="#"className={btnClass}>{btnName}</a>;
  ```

## Element
#### Create An Element
```js
const btnElement = <ahref="#"className="btn btn-primary">Buy</a>;
```

#### Render An Element
Render a React element into the Document Object Model (DOM).
```
const btnElement = <ahref="#"className="btn btn-primary">Buy</a>;
ReactDOM.render(btnElement,document.getElementById('root'));
```

## Component
#### Concepts 
- Written as a JavaScript class or a function.
- Composed of
   - Elements
   - Props
   - State

#### Create A Component
- Key Points
   - The class inherits the `React.Component`.
   - The component name should start with an uppercase letter.
   - The class has `render()` method to return the React element.
- Examples
  ```js
  import React from 'react';
  class YourComponentName extends React.Component {
      render() {
          // some operations
          return (
              // React element
          );
      }
  }
  ```
- Render Method
    - `render()` method needs to return a single React element.
      ```js
      render() {
          return (
              <div>              // div wraps 3 card elements as one element.
                  <card></card>
                  <card></card>
                  <card></card>
              </div>
          );
      }
      ```
  
#### Use A Component
Use it as a DOM tag in JSX. For example, if your component name is `Card`, you can use it as `<Card/>`.
```js
// Define Card component
class Card extends React.Component {
    // some functions
    
    render() {
        return (
        )
    }
}

// Use Card component
reactDOM.render(<Card/>, document.getElementById('root'));
```

#### Import A Component

## Props
#### Concepts 
- Pass of data from parent component to child components.
