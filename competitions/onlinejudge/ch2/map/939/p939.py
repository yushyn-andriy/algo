import sys
from enum import Enum



stdin = sys.stdin


class GenomeType(Enum):
    dominant = 'dominant'
    non_existent = 'non-existent'
    recessive = 'recessive'


    @classmethod
    @property
    def types(cls):
        return [getattr(cls, name).value for name in cls.__dict__['_member_names_']]


class Person:
    __slots__ = ['name', 'genome_type', 'parents', 'parent_names']

    def __init__(self, name, genome_type, parent_names) -> None:
        self.name = name
        self.genome_type = genome_type or None
        self.parents = []
        self.parent_names = parent_names or []

    def __repr__(self) -> str:
        return f'P({self.name}, {self.genome_type}, {self.parents})'



def determine_genome_type(name, persons, cache):
    if name in cache:
        return cache[name]
    
    person = persons[name]
    if person.genome_type:
        return person.genome_type
    else:
        p1, p2 = person.parents

        if not p1.genome_type:
            p1.genome_type = determine_genome_type(p1.name, persons, cache)
        if not p2.genome_type:
            p2.genome_type = determine_genome_type(p2.name, persons, cache)

        if p1.genome_type == p2.genome_type:
            person.genome_type = p1.genome_type 
        elif p1.genome_type == GenomeType.dominant or p2.genome_type == GenomeType.dominant:
            if p1.genome_type == GenomeType.dominant and p2.genome_type == GenomeType.dominant:
                person.genome_type = p1.genome_type 
            elif (
                p1.genome_type == GenomeType.dominant and 
                p2.genome_type == GenomeType.recessive
            ) or (
                p2.genome_type == GenomeType.dominant and 
                p1.genome_type == GenomeType.recessive
            ):
                person.genome_type = GenomeType.dominant
            elif (
                    p1.genome_type == GenomeType.dominant and 
                    p2.genome_type == GenomeType.non_existent
                ) or (
                    p2.genome_type == GenomeType.dominant and 
                    p1.genome_type == GenomeType.non_existent
            ):
                    person.genome_type = GenomeType.recessive
            else:
                if p1.genome_type != GenomeType.dominant:
                    person.genome_type = p1.genome_type 
                else:
                    person.genome_type = p2.genome_type
        else:
            person.genome_type = GenomeType.non_existent
            
    cache[person.name] = person.genome_type
    return person.genome_type
    


if __name__ == '__main__':
    t = int(stdin.readline().strip())
    persons = {}
    for _ in range(t):
        arg1, arg2 = stdin.readline().strip().split()
        
        if arg2 in GenomeType.types:
            child_name, genome_type = arg1, getattr(GenomeType, arg2.replace('-', '_'))
            if child_name not in persons:
                person = Person(child_name, genome_type, [])
                persons[child_name] = person
            else:
                persons[child_name].genome_type= genome_type
        else:
            parent_name, child_name,  = arg1, arg2
            if child_name not in persons:
                person = Person(child_name, None, [parent_name])
                persons[child_name] = person
            else:
                persons[child_name].parent_names.append(parent_name)

    for person in persons.values():
        for parent_name in person.parent_names:
            person.parents.append(persons[parent_name]) 


    cache = {}
    for name in sorted(persons.keys()):
        print(name, determine_genome_type(name, persons, cache).value)

