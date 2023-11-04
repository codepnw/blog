import React from "react"
import { Container, Row, Col } from "react-bootstrap"
import { useParams, useNavigate } from "react-router-dom"
import axios from "axios"

const Delete = () => {
  const params = useParams()
  const navigate = useNavigate()

  const handleDelete = async () => {
    try {
      const apiUrl = import.meta.env.VITE_APP_API + "/" + params.id
      const response = await axios.delete(apiUrl)

      if (response.status === 200) {
        navigate("/", {
          state: "Record deleted successfully.",
        })
      }
    } catch (error) {
      console.log(error)
    }
  }

  return (
    <>
      <Container>
        <Row>
          <h1>Are you sure to delete this record?</h1>
          <Col xs="12" className="py-5">
            <button className="btn btn-danger py-2 my-2" onClick={handleDelete}>
              Procees
            </button>
          </Col>
        </Row>
      </Container>
    </>
  )
}

export default Delete
