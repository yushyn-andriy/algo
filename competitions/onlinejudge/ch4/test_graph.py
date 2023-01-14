from pytest import fixture

from  graph import Graph, Vertex, Edge


@fixture
def directed_graph():
    return Graph(directed=True)


@fixture
def vertex1():
    return Vertex(1)


@fixture
def vertex2():
    return Vertex(2)


def test_create_directed_graph(directed_graph):
    g = directed_graph
    assert g.is_directed() is True
    assert g._incoming is not g._outgoing


def test_insert_vertex(directed_graph, vertex1, vertex2):
    directed_graph.insert_vertex(vertex1)
    directed_graph.insert_vertex(vertex2)
    assert len(directed_graph.vertices()) == 2


def test_insert_edge_one_direction(directed_graph, vertex1, vertex2):
    directed_graph.insert_vertex(vertex1)
    directed_graph.insert_vertex(vertex2)
    directed_graph.insert_edge(vertex1, vertex2)
    assert len(directed_graph.edges()) == 1


def test_insert_edge_both_direction(directed_graph, vertex1, vertex2):
    directed_graph.insert_vertex(vertex1)
    directed_graph.insert_vertex(vertex2)
    directed_graph.insert_edge(vertex1, vertex2)
    directed_graph.insert_edge(vertex2, vertex1)
    assert len(directed_graph.edges()) == 2
