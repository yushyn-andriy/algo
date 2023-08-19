from sqlalchemy import (
    Table,
    Column,
    Integer,
    String,
    Date,
    ForeignKey,
    create_engine,
    text,
)
from sqlalchemy import and_, or_
from sqlalchemy.orm import registry, relationship, sessionmaker
from pathlib import Path



from sqlalchemy import create_engine, MetaData
from sqlalchemy.ext.declarative import declarative_base

import time


DATABASE_URL = 'postgresql://web:123@localhost:5432/dvdrental'
engine = create_engine(DATABASE_URL, echo=True)  # echo=True will print SQL statements.


mapper_registry = registry()
Base = mapper_registry.generate_base()

metadata = mapper_registry.metadata
metadata.reflect(bind=engine)



customer_table = Table(
    'customer',
    metadata,
    autoload=True,
    autoload_with=engine,
)


payment_table = Table(
    'payment',
    metadata,
    autoload=True,
    autoload_with=engine,
)

class Customer(Base):
    __table__ = customer_table

    payments = relationship('Payment', back_populates='customer')

class Payment(Base):
    __table__ = payment_table


    customer = relationship(Customer, back_populates='payments')


from sqlalchemy.orm import joinedload
import pprint


if __name__ == "__main__":
    # engine = create_engine(DATABASE_URL, echo=True, pool_size=10)
    session = sessionmaker(bind=engine)()

    q = '''
SELECT * FROM customer ORDER BY customer_id LIMIT :limit;
'''
    # r = session.execute(text(q), params={'limit': 10}).fetchall()
    # print(type(r), r)

    # pprint.pprint(list(Customer.__table__.columns.items()))
    # r1, *r2 = session.query(Customer).join(Payment).order_by('customer_id').limit(1).all()
    # print(r1.payments[:10])

    r1, *r2 = (session.query(Customer, Payment).join(Payment)
            .filter(and_(
                Customer.customer_id == 1,
                Payment.payment_id == 18495
            ))
            .order_by(Customer.customer_id)
           .limit(1)
           .all())
    print(r1[1].payment_id, r1[0].customer_id)
    # print(r1.payments[:10])

    # time.sleep(10)

