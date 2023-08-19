-- JOIN


-- How can you produce a list of the start times for bookings by members named 'David Farrell'?

select b.starttime from cd.members m inner join cd.bookings b
on b.memid = m.memid
where m.firstname || ' ' || m.surname = 'David Farrell';


-- Question
-- How can you produce a list of the start times for bookings for tennis courts, 
-- for the date '2012-09-21'? Return a list of start time and facility name pairings, ordered by the time.

select b.starttime, f.name from 
cd.bookings b
inner join cd.facilities f on
f.facid = b.facid
where cast (b.starttime as date) = '2012-09-21'
and f.name LIKE 'Tennis %'
order by b.starttime;


-- Question
-- How can you output a list of all members who have recommended another member? 
-- Ensure that there are no duplicates in the list, and that results are ordered by (surname, firstname).

select distinct m.firstname, m.surname
from cd.members m inner join cd.members m1
on m.memid = m1.recommendedby
order by m.surname, m.firstname;




-- How can you output a list of all members, including 
-- the individual who recommended them (if any)? 
-- Ensure that results are ordered by (surname, firstname).

select m.firstname as memfname, m.surname as memsname, 
m1.firstname as recfname, m1.surname as recsname
	from cd.members m 
		left outer join cd.members m1
			on m1.memid = m.recommendedby
order by memsname, memfname;


-- How can you produce a list of all members who
--  have used a tennis court? Include in your output the name 
--  of the court, and the name of the member formatted as 
-- a single column. Ensure no duplicate data, and order by 
-- the member name followed by the facility name.

select distinct m.firstname || ' ' || m.surname as member, f.name as facility
from cd.members m
inner join cd.bookings b
on b.memid = m.memid
inner join cd.facilities f
on f.facid = b.facid
where f.name in ('Tennis Court 1', 'Tennis Court 2')
order by member, facility;




-- How can you produce a list of bookings on the 
-- day of 2012-09-14 which will cost the member (or guest) 
-- more than $30? Remember that guests have different costs to 
-- members (the listed costs are per half-hour 'slot'), and the guest 
-- user is always ID 0. Include in your output the name of the facility, the name of the member
--  formatted as a single column, and the cost. Order by descending cost, and do not use any subqueries.

select m.firstname || ' ' || m.surname as member, max(b.slots * f.membercost) as cost

from cd.members m
inner join cd.bookings b
on b.memid = b.memid
inner join cd.facilities f
on f.facid = b.facid
where cast (b.starttime as DATE) = '2012-09-14'
group by member
HAVING
	MAX (b.slots * f.membercost) > 30
order by cost desc;

--

-- How can you produce a list of bookings on the 
-- day of 2012-09-14 which will cost the member (or guest) 
-- more than $30? Remember that guests have different costs to 
-- members (the listed costs are per half-hour 'slot'), and the guest 
-- user is always ID 0. Include in your output the name of the facility, the name of the member
--  formatted as a single column, and the cost. Order by descending cost, and do not use any subqueries.
select 
    m.firstname || ' ' || m.surname as member,
	f.name as facility,
	case
		when m.memid = 0 then b.slots * f.guestcost
		when m.memid <> 0 then b.slots * f.membercost
	end as cost
from 
	cd.members m
	inner join cd.bookings b
		on b.memid = m.memid
	inner join cd.facilities f
		on f.facid = b.facid
where
	cast (b.starttime as DATE) = '2012-09-14' and
	case
		when m.memid != 0 then b.slots * f.membercost > 30
		when m.memid = 0 then b.slots * f.guestcost > 30
	end 
order by cost desc;

--

select member, facility, cost from (
	select 
		mems.firstname || ' ' || mems.surname as member,
		facs.name as facility,
		case
			when mems.memid = 0 then
				bks.slots*facs.guestcost
			else
				bks.slots*facs.membercost
		end as cost
		from
			cd.members mems
			inner join cd.bookings bks
				on mems.memid = bks.memid
			inner join cd.facilities facs
				on bks.facid = facs.facid
		where
			bks.starttime >= '2012-09-14' and
			bks.starttime < '2012-09-15'
	) as bookings
	where cost > 30
order by cost desc; 

--
select mems.firstname || ' ' || mems.surname as member, 
	facs.name as facility, 
	case 
		when mems.memid = 0 then
			bks.slots*facs.guestcost
		else
			bks.slots*facs.membercost
	end as cost
        from
                cd.members mems                
                inner join cd.bookings bks
                        on mems.memid = bks.memid
                inner join cd.facilities facs
                        on bks.facid = facs.facid
        where
		bks.starttime >= '2012-09-14' and 
		bks.starttime < '2012-09-15' and (
			(mems.memid = 0 and bks.slots*facs.guestcost > 30) or
			(mems.memid != 0 and bks.slots*facs.membercost > 30)
		)
order by cost desc;  


-- How can you output a list of all members, 
-- including the individual who recommended them (if any), 
-- without using any joins? Ensure that there are no 
-- duplicates in the list, and that each firstname + surname pairing 
-- is formatted as a column and ordered.

select distinct m.firstname || ' ' || m.surname as member,
case 
	when m.recommendedby is null then null
	else
		(select cd.members.firstname || ' ' || cd.members.surname from cd.members
		where cd.members.memid = m.recommendedby)
end recommender
from cd.members m
order by member;



-- ============================================
-- How can you retrieve all the information from the cd.facilities table?
select * from cd.facilities;


-- You want to print out a list of all of the facilities 
-- and their cost to members. How would you retrieve 
-- a list of only facility names and costs?
-- select f.name, f.membercost from cd.facilities f;

select f.name, f.membercost from cd.facilities f;


-- How can you produce a list of facilities that charge a fee to members?
select * from cd.facilities f where f.membercost>0;


-- How can you produce a list of facilities 
-- that charge a fee to members, and that fee 
-- is less than 1/50th of the monthly maintenance cost? 
-- Return the facid, facility name, member cost, 
-- and monthly maintenance of the facilities in question.

select f.facid, f.name, f.membercost, f.monthlymaintenance from cd.facilities f 
where f.membercost < f.monthlymaintenance/50 and f.membercost>0 ;


-- How can you produce a list of all facilities with the word 'Tennis' in their name?
select * from cd.facilities f where
f.name like '%Tennis%';


-- How can you retrieve the details of facilities with 
-- ID 1 and 5? Try to do it without using the OR operator.
select * from cd.facilities f where f.facid in (1, 5);




-- How can you produce a list of facilities, 
-- with each labelled as 'cheap' or 'expensive' depending on if their monthly 
-- maintenance cost is more than $100? 
-- Return the name and monthly maintenance of the facilities in question.

select f.name, 
case
when f.monthlymaintenance > 100  then 'expensive'
else
	'cheap'
end cost from cd.facilities f;



-- How can you produce a list of members who joined after the start 
-- of September 2012? Return the memid, surname, firstname, 
-- and joindate of the members in question.

select m.memid, m.surname, m.firstname, m.joindate from cd.members m
where 
	cast (m.joindate as date) >= '2012.09.01';


-- How can you produce an ordered list of 
-- the first 10 surnames in the members table? 
-- The list must not contain duplicates.
select distinct m.surname from cd.members m
order by m.surname
limit 10;


-- You, for some reason, want a combined list of 
-- all surnames and all facility names. 
-- Yes, this is a contrived example :-). Produce that list!
select f.name from cd.facilities f
union
select m.surname from cd.members m;



-- You'd like to get the signup date of your last member. How can you retrieve this information?
select m.joindate as latest from cd.members m
order by m.joindate desc limit 1;
select max(m.joindate) as latest from cd.members m;


select m.firstname, m.surname, m.joindate from cd.members m



order by m.joindate desc limit 1;
select m.firstname, m.surname, m.joindate from  cd.members m
where m.joindate = (select max(ms.joindate) from cd.members ms);


-- MODIFYING DATA

-- The club is adding a new facility - a spa. 
-- We need to add it into the facilities table. Use the following values:

-- facid: 9, Name: 'Spa', membercost: 20, guestcost: 30, initialoutlay: 100000,
--  monthlymaintenance: 800.
insert into cd.facilities
values
(9, 'Spa', 20, 30, 100000, 800);

insert into cd.facilities
values
(9, 'Spa', 20, 30, 100000, 800),
(10, 'Squash Court 2', 3.5, 17.5, 5000, 80);




-- Let's try adding the spa to the facilities table again. This time, though, we want to automatically generate the value for the next facid, rather than specifying it as a constant. Use the following values for everything else:

-- Name: 'Spa', membercost: 20, guestcost: 30, initialoutlay: 100000, monthlymaintenance: 800.
insert into cd.facilities(facid, name, membercost, guestcost, initialoutlay, monthlymaintenance)
values
((select max(facid) + 1 from cd.facilities), 'Spa', 20, 30, 100000, 800);


update cd.facilities set
initialoutlay = 10000
where facid = 1;


-- We want to increase the price of the tennis courts for 
-- both members and guests. Update the costs to be 6 for members, and 30 for guests.

update cd.facilities set
membercost = 6,
guestcost = 30
where name LIKE 'Tennis%';


-- We want to alter the price of the second tennis court so that it 
-- costs 10% more than the first one. Try to do this
--  without using constant values for the prices, 
--  so that we can reuse the statement if we want to.

update cd.facilities set
membercost = membercost + membercost * 0.1,
guestcost = guestcost + guestcost * 0.1
where name LIKE 'Tennis Court 2';



-- As part of a clearout of our database, we want 
-- to delete all bookings from the cd.bookings table. How can we accomplish this?
delete from cd.bookings;

-- We want to remove member 37, who has never made a booking, 
-- from our database. How can we achieve that?
delete from cd.members m
where m.memid = 37;



-- In our previous exercises, we deleted a specific 
-- member who had never made a booking. 
-- How can we make that more general, to delete all 
-- members who have never made a booking?
delete from cd.members m
where m.memid not in (select distinct b.memid from cd.bookings b);


--  AGGREGATES

-- For our first foray into aggregates, we're going to stick to something simple.
-- We want to know how many facilities exist - simply produce a total count.

select count(*) from cd.facilities;



-- Produce a count of the number of facilities that have a cost to guests of 10 or more.
select count(*) from cd.facilities f
where f.guestcost > 10;



select m.recommendedby, count(*) from cd.members m 
where m.memid != 0 and m.recommendedby is not null
group by m.recommendedby
order by m.recommendedby;


-- Produce a list of the total number of slots booked per facility. 
-- For now, just produce an output table 
-- consisting of facility id and slots, sorted by facility id.
select 
	b.facid,
	sum(b.slots)
from cd.bookings b
group by b.facid
order by b.facid;


-- Produce a list of the total number of slots booked per facility in the month 
-- of September 2012. Produce an output table consisting 
-- of facility id and slots, sorted by the number of slots.
select b.facid, sum(b.slots) as "Total Slots" from cd.bookings b
where 
	cast (b.starttime as date) 
	between '2012.09.01' and '2012.10.01'
group by b.facid
order by "Total Slots";



-- Produce a list of the total number of slots
--  booked per facility per month in the year 
-- of 2012. Produce an output table consisting of 
-- facility id and slots, sorted by the id and month.
select b.facid, extract(month from b.starttime) as month, sum(b.slots) as "Total Slots" from cd.bookings b
where extract(year from b.starttime) = '2012'
group by (b.facid, month)
order by b.facid, month;



-- Find the total number of members (including guests) 
-- who have made at least one booking.
select count(al.qcount) from (select count(*) as qcount 
  from cd.bookings b group by b.memid) al;


-- Produce a list of facilities with more than 1000 slots booked. 
-- Produce an output table consisting of facility id and slots, sorted by facility id.
select ag.facid, ag.sum as "Total Slots" from (select facid, sum(slots) from cd.bookings
group by facid
order by facid) ag
where ag.sum > 1000;

select facid, sum(slots) as "Total Slots"
        from cd.bookings
        group by facid
        having sum(slots) > 1000
        order by facid 

-- Produce a list of facilities along with their total revenue. 
-- The output table should consist of facility name 
-- and revenue, sorted by revenue. 
-- Remember that there's a different cost for guests and members!

select f.name, sum(
	case
  		when b.memid <> 0 then b.slots * f.membercost
  		else
  			b.slots * f.guestcost
	end   
) revenue from cd.bookings b inner join cd.facilities f on
b.facid = f.facid
group by f.name
order by revenue;



-- Produce a list of facilities with a total revenue less than 1000. 
-- Produce an output table consisting of facility 
-- name and revenue, sorted by revenue. 
-- Remember that there's a different cost for guests and members!

select f.name, sum(
	case
  		when b.memid <> 0 then b.slots * f.membercost
  		else
  			b.slots * f.guestcost
	end   
) revenue from cd.bookings b inner join cd.facilities f on
b.facid = f.facid
group by f.name
having sum(
	case
  		when b.memid <> 0 then b.slots * f.membercost
  		else
  			b.slots * f.guestcost
	end   
) < 1000
order by revenue;

select * from (select f.name, sum(
	case
  		when b.memid <> 0 then b.slots * f.membercost
  		else
  			b.slots * f.guestcost
	end   
) revenue from cd.bookings b inner join cd.facilities f on
b.facid = f.facid
group by f.name
order by revenue) ag
where ag.revenue < 1000;


--  Output the facility id that has the highest number of slots booked.
--  For bonus points, try a version without a LIMIT clause. 
--  This version will probably look messy!

select b1.facid, sum(b1.slots) as total from cd.bookings b1
  	group by b1.facid
	having sum(b1.slots) = (
			select max(ag.total) 
	  		from (
			  select b.facid, sum(b.slots) as total 
			  from cd.bookings b
  			  group by b.facid) ag
	  		);

with sum as (select facid, sum(slots) as totalslots
	from cd.bookings
	group by facid
)
select facid, totalslots 
	from sum
	where totalslots = (select max(totalslots) from sum);