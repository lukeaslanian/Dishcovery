 import React, { useState, useEffect } from 'react';
 import { getRecipes, createRecipe, updateRecipe, deleteRecipe } from './api';
 import './App.css';

 function App() {
     const [recipes, setRecipes] = useState([]);
     const [title, setTitle] = useState('');
     const [description, setDescription] = useState('');

     useEffect(() => {
         fetchRecipes();
     }, []);

     const fetchRecipes = async () => {
         const result = await getRecipes();
         setRecipes(result.data);
     };

     const handleAddRecipe = async () => {
         const newRecipe = { title, description };
         await createRecipe(newRecipe);
         fetchRecipes();
         setTitle('');
         setDescription('');
     };

     const handleUpdateRecipe = async (id) => {
         const updatedRecipe = { title, description };
         await updateRecipe(id, updatedRecipe);
         fetchRecipes();
     };

     const handleDeleteRecipe = async (id) => {
         await deleteRecipe(id);
         fetchRecipes();
     };

     return (
         <div className="App">
             <h1>Recipes</h1>
             <input
                 type="text"
                 placeholder="Title"
                 value={title}
                 onChange={(e) => setTitle(e.target.value)}
             />
             <input
                 type="text"
                 placeholder="Description"
                 value={description}
                 onChange={(e) => setDescription(e.target.value)}
             />
             <button onClick={handleAddRecipe}>Add Recipe</button>
             <ul>
                 {recipes.map((recipe) => (
                     <li key={recipe.id}>
                         <h2>{recipe.title}</h2>
                         <p>{recipe.description}</p>
                         <button onClick={() => handleUpdateRecipe(recipe.id)}>Update</button>
                         <button onClick={() => handleDeleteRecipe(recipe.id)}>Delete</button>
                     </li>
                 ))}
             </ul>
         </div>
     );
 }

 export default App;
 ```
