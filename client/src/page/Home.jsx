import React from "react"
import { useEffect, useState } from "react"
import { Container, Row, Col } from "react-bootstrap"
import { Link } from "react-router-dom"
import Spinner from "react-bootstrap/Spinner"
import axios from "axios"

const Home = () => {
  const [apiData, setApiData] = useState(false)
  const [loading, setLoading] = useState(true)

  useEffect(() => {
    const fetchData = async () => {
      try {
        const apiUrl = import.meta.env.VITE_APP_API
        const response = await axios.get(apiUrl)

        if (response.status === 200) {
          if (response.data?.statusText === "OK") {
            setApiData(response?.data?.blog_records)
          }
        }

        setLoading(false)
      } catch (error) {
        setLoading(false)
        console.log(error.response)
      }
    }

    fetchData()
    return () => {}
  }, [])

  console.log(apiData)

  if (loading) {
    return (
      <Container className="spinner">
        {" "}
        <Spinner animation="grow" />{" "}
      </Container>
    )
  }

  return (
    <Container>
      <Row>
        <Col xs="12" className="text-center py-2">
          <Link to="/add" className="btn btn-primary"> Add Post</Link>
        </Col>

        {apiData &&
          apiData.map((record, index) => (
            <Col key={index} xs="4" className="py-5 box">
              <div className="title">
                <Link to={`/blog/${record.id}`}>{record.title}</Link>
              </div>
              <div> {record.post} </div>
            </Col>
          ))}
      </Row>
    </Container>
  )
}

export default Home
