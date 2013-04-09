Outages
#######


When outages happen (and they will) it is incredibly useful to have a point person that will handle communcations while the admin focuses on resolving the issue.   A common communication forum (conference call, irc, etc) that affected users can jump in get a quick status and pass along any information is useful.  On a larger scale you would want to have a status page that is constantly updated.

Once an outage takes place as an Admin you have to primary functions:

1.  Restore service.
2.  Document what caused the failure.

Depending on the failure and the impact #2 may actually be more important than #1.   Assuming you have a proper DEV / QA / PROD / DR enviroment setup you should fail over to your DR systems and take your time investigating the PROD systems to identify root cause (and build a test so you can replicate the problem at will) and come up with a proper resolution that can be documented and isn't considered a "band-aide" fix.    Apply the change up through DEV & QA and the then to your PROD, assuming all goes well fail back from DR, run your replicated failure test to show PROD isn't going to fail from that issue again.   Then you are safe to make the change on your DR side.

Post outage you need to document what the problem was, what the root cause was, and what the steps were to fix.   Depending on who your users are you may wish to provide this information to them as well to keep the informed.

