import axios from "axios";
import { useEffect, useState } from "react";
import { Badge, Button, Card, Container, Form } from "react-bootstrap";
import { RiArrowDownSLine } from "react-icons/ri";
import "./App.css";

function App() {
  const [endNode, setEndNode] = useState(".45 GVT Armor Piercing");
  const [list, setList] = useState([]);
  const [path, setPath] = useState([]);

  useEffect(() => {
    axios.get("http://localhost:8080/getAdjList").then((response) => {
      setList(response.data);
    });
  }, []);

  const findProgress = () => {
    axios
      .get("http://localhost:8080/findPath", {
        params: { endNode: endNode },
      })
      .then((response) => {
        console.log(response.data);
        setPath(response.data);
      });
  };

  return (
    <div className="App">
      <Container style={{ padding: "50px" }}>
        <h1>MODERN WARFARE 2 PATHFINDER</h1>
        <Form.Select
          style={{ width: "50%", margin: "auto", marginTop: "50px" }}
          value={endNode}
          onChange={(e) => {
            setEndNode(e.currentTarget.value);
          }}
          aria-label="Default select example"
        >
          {list &&
            list.map((item, id) => (
              <option key={id} value={item}>
                {item}
              </option>
            ))}
        </Form.Select>

        <Button
          style={{ marginTop: "50px" }}
          variant="outline-success"
          onClick={findProgress}
        >
          FIND PROGRESSION
        </Button>

        <Container style={{ padding: "50px" }}>
          {path &&
            path.map((item, id) => (
              <div key={id}>
                <Card
                  style={{
                    width: "100%",
                    background: "#333",
                    marginBottom: "20px",
                  }}
                >
                  <Card.Body>
                    <Card.Title>
                      {id !== 0 && "unlocks"} {item.object}{" "}
                      <Badge bg="dark" text="success">
                        {item.type}
                      </Badge>
                    </Card.Title>
                  </Card.Body>
                </Card>
                {id !== path.length - 1 && (
                  <h5
                    className="mb-2"
                    style={{
                      color: "#bbb",
                    }}
                  >
                    {id !== path.length - 1 && "at level " + item.level}
                    <RiArrowDownSLine
                      color="#198754"
                      size="1.5em"
                      style={{ marginLeft: "10px" }}
                    />
                  </h5>
                )}
              </div>
            ))}
        </Container>
      </Container>
    </div>
  );
}

export default App;
