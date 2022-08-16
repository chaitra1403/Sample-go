from boto3.session import Session

AKEY='paste-your-aws-Access-KEy'
SKEY='paste-your-aws-secret-Key'

session = Session(aws_access_key_id=AKEY,
                  aws_secret_access_key=SKEY)
s3 = session.resource('s3')
bkt = s3.Bucket('roost.ai') #give the s3  bucket name

for i in bkt.objects.all():
    print(i.key)
