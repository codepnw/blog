import { Routes, Route } from "react-router-dom"

import "./App.css"
import Home from "./page/Home"
import Blog from "./page/Blog"
import Header from "./components/layout/Header"
import Footer from "./components/layout/Footer"
import Add from "./page/Add"
import Edit from "./page/Edit"
import Delete from "./page/Delete"

const App = () => {
  return (
    <>
      <Header />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/add" element={<Add />} />
        <Route path="/edit/:id" element={<Edit />} />
        <Route path="/blog/:id" element={<Blog />} />
        <Route path="/delete/:id" element={<Delete />} />
      </Routes>
      <Footer />
    </>
  )
}

export default App
