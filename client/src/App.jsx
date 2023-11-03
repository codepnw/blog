import { Routes, Route } from "react-router-dom"

import "./App.css"
import Home from "./page/Home"
import Blog from "./page/Blog"
import Header from "./components/layout/Header"
import Footer from "./components/layout/Footer"
import Add from "./page/Add"

const App = () => {
  return (
    <>
      <Header />
      <Routes>
        <Route path="/" element={<Home />} />
        <Route path="/blog/:id" element={<Blog />} />
        <Route path="/add" element={<Add />} />
      </Routes>
      <Footer />
    </>
  )
}

export default App
