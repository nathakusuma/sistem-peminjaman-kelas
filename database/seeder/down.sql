DELETE
FROM replies
WHERE id IN (
             '019750be-61c6-7344-a711-d491af4dd71d',
             '019750be-61c6-7344-a711-d491af4dd71e',
             '019750be-61c6-7344-a711-d491af4dd71f',
             '019750be-61c6-7344-a711-d491af4dd720',
             '019750be-61c6-7344-a711-d491af4dd721',
             '019750be-61c6-7344-a711-d491af4dd722',
             '019750be-61c6-7344-a711-d491af4dd723',
             '019750be-61c6-7344-a711-d491af4dd724',
             '019750be-61c6-7344-a711-d491af4dd725',
             '019750be-61c6-7344-a711-d491af4dd726',
             '019750be-61c6-7344-a711-d491af4dd727',
             '019750be-61c6-7344-a711-d491af4dd728',
             '019750be-61c6-7344-a711-d491af4dd729',
             '019750be-61c6-7344-a711-d491af4dd72a',
             '019750be-61c6-7344-a711-d491af4dd72b'
    );

DELETE
FROM proposals
WHERE id IN (
             '019750be-61c6-7344-a711-d491af4dd71d',
             '019750be-61c6-7344-a711-d491af4dd71e',
             '019750be-61c6-7344-a711-d491af4dd71f',
             '019750be-61c6-7344-a711-d491af4dd720',
             '019750be-61c6-7344-a711-d491af4dd721',
             '019750be-61c6-7344-a711-d491af4dd722',
             '019750be-61c6-7344-a711-d491af4dd723',
             '019750be-61c6-7344-a711-d491af4dd724',
             '019750be-61c6-7344-a711-d491af4dd725',
             '019750be-61c6-7344-a711-d491af4dd726',
             '019750be-61c6-7344-a711-d491af4dd727',
             '019750be-61c6-7344-a711-d491af4dd728',
             '019750be-61c6-7344-a711-d491af4dd729',
             '019750be-61c6-7344-a711-d491af4dd72a',
             '019750be-61c6-7344-a711-d491af4dd72b',
             '019750be-61c6-7344-a711-d491af4dd72c',
             '019750be-61c6-7344-a711-d491af4dd72d',
             '019750be-61c6-7344-a711-d491af4dd72e',
             '019750be-61c6-7344-a711-d491af4dd72f',
             '019750be-61c6-7344-a711-d491af4dd730'
    );

DELETE
FROM users
WHERE email IN (
                'student1@example.com',
                'student2@example.com',
                'student3@example.com',
                'student4@example.com',
                'student5@example.com',
                'admin1@example.com',
                'admin2@example.com'
    );
