# React

## Installation
- Create React App
  ```bash
  npm install -g create-react-app
  create-react-app <your app name>
  ```

## JSX
### Concepts 
- Syntax extension to JavaScript.
- JSX produces React “elements”.

### Differences between JSX and HTML
- JSX use “ClassName” rather than “Class” in HTML to define a class name.
- JSX use the camel-case naming convention.

### Example
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
### Create An Element
```js
const btnElement = <ahref="#"className="btn btn-primary">Buy</a>;
```

### Render
Render a React element into the Document Object Model (DOM).
```
const btnElement = <ahref="#"className="btn btn-primary">Buy</a>;
ReactDOM.render(btnElement,document.getElementById('root'));
```

## Component

