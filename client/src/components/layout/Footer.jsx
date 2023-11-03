import React from "react"
import { Container, Row, Col } from "react-bootstrap"
import { Link } from "react-router-dom"

const Footer = () => {
  return (
    <Container fluid className="container-fluid footer">
      <Row>
        <Col xs="6">
          <ul>
            <li>
              <Link to="/">Home</Link>
            </li>
            <li>
              <Link to="/">Blog</Link>
            </li>
            <li>
              <Link to="/">About</Link>
            </li>
            <li>
              <Link to="/">Contact</Link>
            </li>
          </ul>
        </Col>
        <Col xs="6">
          <div className="col text-center text-lg-right pt-3"></div>
          <a href="#">
            <h3>Connect With Us</h3>
          </a>
        </Col>
      </Row>
    </Container>
  )
}

export default Footer
