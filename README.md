# GehtGüd - A Software Engineer Interview Support Suite
GehtGüd is a (in-progress) collection of tools intended to support the sotware engineer in improving their job hunting and interview skills. It is intended to do everything it can to help you improve, track progress and get good at interviewing (everything shy of putting in the work).  This is the tooling that I am personally using for tracking my progress so expect it to grow and change overtime (as I identify various features that I need/want).

This application (for now) exists entirely in the command line but may have some TUI/local-web features in the near future.

# The (Planned) Services:
### Leetcode Tracker - A tracker for all attempts on Leetcode interviews and problems.
- Log problems you have encountered
- Track you attempts
- Store your solutions
- Review previous attempts with spaced repetition for maximum learning efficiency
- Tracks where you have seen (ie. which companies) certain problems and when
- Create study lists and track your progress
- Get daily recomendations per your plan

### Resume, Application and Interview Tracker
- A datastore for all applications sent
  - What company
  - What resume you sent
  - Job Title
  - Job Description
  - Did you get an interview
  - What stage did are you on
  - What was the offer (if any)
  - Notes/feedback impressions

### Resume Generator
- Give a locally run llm your resume and previous roles and the job you are applying for
- Let it tell you if it is a good fit
- Have it generate a resume the ideal resume for that job (without lying)

### Mock Interview Generator
- Chooses two LC questions from a pool of Mediums and Hards and One System Design Question
- Gives you a suggested timeframe
- Lets you track your mock interviews over time
- Mark which questions you succeeded and which you failed and why
- Saves your attempt and adds it to the spaced repetition.

### Habit Tracker
- Give it a range and it will return, by day, your "events" (Applications sent and LC problems solved)
- Can also just return a percentage based on your goal (Goal is 2 LC a day it will give you a percentage for each day)


