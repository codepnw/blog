import React, { useEffect, useState } from "react"
import { Container, Col, Row } from "react-bootstrap"
import { useParams } from "react-router-dom"
import axios from "axios"

const Blog = () => {
  const params = useParams()
  const [apiData, setApiData] = useState(false)
  useEffect(() => {
    const fetchData = async () => {
      try {
        const apiUrl = import.meta.env.VITE_APP_API + "/" + params.id

        const response = await axios.get(apiUrl)

        if (response.status === 200) {
          if (response.data?.statusText === "OK") {
            setApiData(response?.data?.record)
          }
        }
      } catch (error) {
        console.log(error.response)
      }
    }

    fetchData()
    return () => {}
  }, [])

  console.log(apiData)
  return (
    <Container>
      <Row>
        <Col xs="12">
          <h2> {apiData.title} </h2>
        </Col>
        <Col xs="12">
          <p> {apiData.post} </p>
        </Col>
      </Row>
    </Container>
  )
}

export default Blog
