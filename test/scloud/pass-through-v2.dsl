
events = read-splunk-firehose();
write-index(events, "index", "main");
