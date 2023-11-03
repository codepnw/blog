import { useEffect, useState } from "react"
import { Container, Row, Col } from "react-bootstrap"
import axios from "axios"
import "./App.css"

const App = () => {
  const [apiData, setApiData] = useState(false)
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
      } catch (error) {
        console.log(error.response)
      }
    }

    fetchData()
    return () => {}
  }, [])

  console.log(apiData)

  return (
    <>
      <Container>
        <Row>
          <Col xs="12" className="text-center py-2">
            <h1 className="text-center">Blog App with Go Fiber</h1>
          </Col>

          {apiData &&
            apiData.map((record, index) => (
              <Col key={index} xs="4" className="py-5 box">
                <div className="title"> {record.title} </div>
                <div> {record.post} </div>
              </Col>
            ))}
        </Row>
      </Container>
    </>
  )
}

export default App
