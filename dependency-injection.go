package main

import "fmt"

////////////////////////////
///// The “inner ring” /////
////////////////////////////

// PoemStorage is just an interface that defines the behavior of a poem storage.
// This is all that Poem knows (and needs to know) about storing and retrieving poems.
// Nothing from the “outer ring” appears here.
type PoemStorage interface {
	Type() string        // Return a string describing the storage type.
	Load(string) []byte  // Load a poem by name.
	Save(string, []byte) // Save a poem by name.
}

// A Poem contains some poetry and an abstaract storage reference
type Poem struct {
	content []byte
	storage PoemStorage
}

// NewPoem constructs a Poem object.
// We use this constructor to inject an object
// that satisfies the PoemStorage interface.
func NewPoem(ps PoemStorage) *Poem {
	return &Poem{
		content: []byte("I am a poem from a " + ps.Type() + "."),
		storage: ps,
	}
}

// Save simply calls Save on the interface type.
// The Poem object neither knows nor cares about which actual storage
// object receives this method call.
func (p *Poem) Save(name string) {
	p.storage.Save(name, p.content)
}

// Load also invokes the injected storage object without knowing it.
func (p *Poem) Load(name string) {
	p.content = p.storage.Load(name)
}

// String makes Poem a Stringer, allowing us to drop it anywhere a string would be expected.
func (p *Poem) String() string {
	return string(p.content)
}

////////////////////////////
///// The “outer ring” /////
////////////////////////////

type Notebook struct {
	poems map[string][]byte
}

func NewNotebook() *Notebook {
	return &Notebook{
		poems: map[string][]byte{},
	}
}

// After adding Save and Load, Notebook implicitly satisfies PoemStorage.
func (n *Notebook) Save(name string, contents []byte) {
	n.poems[name] = contents
}

func (n *Notebook) Load(name string) []byte {
	return n.poems[name]
}

// Type returns an informal description of the storage type.
func (n *Notebook) Type() string {
	return "Notebook"
}

// A Napkin is the emergency storage device of a poet. It can store only one poem.
type Napkin struct {
	poem []byte
}

func NewNapkin() *Napkin {
	return &Napkin{
		poem: []byte{},
	}
}

func (n *Napkin) Save(name string, contents []byte) {
	n.poem = contents
}

func (n *Napkin) Load(name string) []byte {
	return n.poem
}

func (n *Napkin) Type() string {
	return "Napkin"
}

func main() {
	// Create and connect objects, then save and load a few poems from different storage objects.
	notebook := NewNotebook()
	napkin := NewNapkin()

	// First, write a poem into a notebook. NewPoem() injects the dependency.
	poem := NewPoem(notebook)
	poem.Save("My first poem")

	// Create a new poem object to prove that the notebook storage works.
	poem = NewPoem(notebook)
	poem.Load("My first poem")
	fmt.Println(poem)

	// Now we do the same with a napkin as storage
	poem = NewPoem(napkin)

	// Note the poem still just uses Save and Load. “Notebook? Napkin? I don’t care.”
	poem.Save("My second poem")
	poem = NewPoem(napkin)
	poem.Load("My second poem")
	fmt.Println(poem)
}
